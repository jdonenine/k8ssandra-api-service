package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jdonenine/k8ssandra-api-service/pkg/resources"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type StatefulSetsController struct {
	Kubeclient    *kubernetes.Clientset
	KubeNamespace string
}

func (controller *StatefulSetsController) GetCassDcStatefulSets(w http.ResponseWriter, r *http.Request) {
	if controller.Kubeclient == nil {
		log.Println("Unable to execute request, no valid client configured.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Parse name from url
	routeParams := mux.Vars(r)
	if routeParams == nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	cassDcName := routeParams["name"]
	if len(cassDcName) < 1 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	resourceManager := resources.ResourceManager{Kubeclient: controller.Kubeclient, KubeNamespace: controller.KubeNamespace}
	_, getCassDcErr := resourceManager.GetCassDc(cassDcName)
	if getCassDcErr != nil {
		if errors.IsNotFound(getCassDcErr) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		log.Printf("Unable to retrieve CassandraDataCenter resource with name '%s' from namesapce '%s', failed with error: '%s'", cassDcName, controller.KubeNamespace, getCassDcErr)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Get the initial list of stateful sets in the namespace
	ssList, ssError := controller.Kubeclient.AppsV1().StatefulSets(controller.KubeNamespace).List(context.TODO(), meta_v1.ListOptions{})
	if ssError != nil {
		log.Printf("Unable to retrieve StatefulSets, failed with error: '%s'", ssError)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Filter down to only the resources owned by the specified cassdc
	filteredItems := []v1.StatefulSet{}
	for _, ss := range ssList.Items {
		match := false
		for _, owner := range ss.OwnerReferences {
			if owner.Kind == "CassandraDatacenter" && owner.Name == cassDcName {
				match = true
				break
			}
		}
		if match {
			filteredItems = append(filteredItems, ss)
		}
	}
	ssList.Items = filteredItems

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	encodeErr := json.NewEncoder(w).Encode(ssList)
	if encodeErr != nil {
		log.Printf("Unable to encode response for listing of statefulset resources from namesapce '%s', failed with error: '%s'", controller.KubeNamespace, encodeErr)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
