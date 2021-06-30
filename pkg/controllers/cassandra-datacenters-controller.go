package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jdonenine/k8ssandra-api-service/pkg/models"
	"github.com/jdonenine/k8ssandra-api-service/pkg/resources"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"
)

type CassandraDatacentersController struct {
	Kubeclient    *kubernetes.Clientset
	KubeNamespace string
}

func (controller *CassandraDatacentersController) GetCassDc(w http.ResponseWriter, r *http.Request) {
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
	cassDc, getCassDcErr := resourceManager.GetCassDc(cassDcName)
	if getCassDcErr != nil {
		if errors.IsNotFound(getCassDcErr) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		log.Printf("Unable to retrieve CassandraDataCenter resource with name '%s' from namesapce '%s', failed with error: '%s'", cassDcName, controller.KubeNamespace, getCassDcErr)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Only return the resource if it's managed by k8ssandra
	appName := cassDc.ObjectMeta.Labels["app.kubernetes.io/name"]
	if appName != "k8ssandra" {
		http.Error(w, "The resource requested is not managed by K8ssandra", http.StatusUnauthorized)
		return
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	encodeErr := json.NewEncoder(w).Encode(cassDc)
	if encodeErr != nil {
		log.Printf("Unable to write response for listing of cassDc resource with name '%s' from namesapce '%s', failed with error: '%s'", cassDcName, controller.KubeNamespace, encodeErr)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (controller *CassandraDatacentersController) GetCassDcs(w http.ResponseWriter, r *http.Request) {
	if controller.Kubeclient == nil {
		log.Println("Unable to execute request, no valid client configured.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	resourceManager := resources.ResourceManager{Kubeclient: controller.Kubeclient, KubeNamespace: controller.KubeNamespace}
	cassDcs, getCassDcsErr := resourceManager.GetCassDcs()
	if getCassDcsErr != nil {
		log.Printf("Unable to retrieve CassandraDatacenterList resource from namesapce '%s', failed with error: '%s'", controller.KubeNamespace, getCassDcsErr)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Filter down to only the resources owned by the specified cassdc
	filteredItems := []models.CassandraDatacenter{}
	for _, cassDc := range cassDcs.Items {
		appName := cassDc.ObjectMeta.Labels["app.kubernetes.io/name"]
		if appName == "k8ssandra" {
			filteredItems = append(filteredItems, cassDc)
		}
	}
	cassDcs.Items = filteredItems

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	encodeErr := json.NewEncoder(w).Encode(cassDcs)
	if encodeErr != nil {
		log.Printf("Unable to write response for listing of CassandraDatacenterList resources from namesapce '%s', failed with error: '%s'", controller.KubeNamespace, encodeErr)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
