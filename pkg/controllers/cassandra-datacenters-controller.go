package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jdonenine/k8ssandra-api-service/pkg/models"
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

	// Query the k8s API for the CassandraDataCenter resource in the given namespace
	resourcePath := fmt.Sprintf("/apis/cassandra.datastax.com/v1beta1/namespaces/%s/cassandradatacenters/%s", controller.KubeNamespace, cassDcName)
	rawCassDc, err := controller.Kubeclient.RESTClient().
		Get().
		AbsPath(resourcePath).
		DoRaw(context.TODO())
	if err != nil {
		log.Printf("Unable to retrieve CassandraDataCenter resource with name '%s' from namesapce '%s', failed with error: '%s'", cassDcName, controller.KubeNamespace, err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Parse the response from the k8s API
	cassDc := models.CassandraDatacenter{}
	parseErr := json.Unmarshal([]byte(rawCassDc), &cassDc)
	if parseErr != nil {
		log.Printf("Unable to parse retrieved CassandraDataCenter resource with name '%s' response from namesapce '%s', failed with error: '%s'", cassDcName, controller.KubeNamespace, parseErr)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	encodeErr := json.NewEncoder(w).Encode(cassDc)
	if encodeErr != nil {
		log.Printf("Unable to encode response for listing of cassDc resource with name '%s' from namesapce '%s', failed with error: '%s'", cassDcName, controller.KubeNamespace, encodeErr)
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

	// Query the k8s API for all of the CassandraDataCenter resources in the given namespace
	resourcePath := fmt.Sprintf("/apis/cassandra.datastax.com/v1beta1/namespaces/%s/cassandradatacenters/", controller.KubeNamespace)
	rawCassdcs, err := controller.Kubeclient.RESTClient().
		Get().
		AbsPath(resourcePath).
		DoRaw(context.TODO())
	if err != nil {
		log.Printf("Unable to retrieve CassandraDataCenter resources from namesapce '%s', failed with error: '%s'", controller.KubeNamespace, err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Parse the response from the k8s API
	cassDcs := models.CassandraDatacenters{}
	parseErr := json.Unmarshal([]byte(rawCassdcs), &cassDcs)
	if parseErr != nil {
		log.Printf("Unable to parse retrieved CassandraDataCenter resource response from namesapce '%s', failed with error: '%s'", controller.KubeNamespace, parseErr)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Build the response object
	responseMetadata := models.ResponseMetadata{Total: len(cassDcs.Items), Count: len(cassDcs.Items)}
	response := struct {
		Items    []models.CassandraDatacenter `json:"items"`
		Metadata models.ResponseMetadata      `json:"metadata"`
	}{
		cassDcs.Items,
		responseMetadata,
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	encodeErr := json.NewEncoder(w).Encode(response)
	if encodeErr != nil {
		log.Printf("Unable to encode response for listing of cassDc resources from namesapce '%s', failed with error: '%s'", controller.KubeNamespace, encodeErr)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
