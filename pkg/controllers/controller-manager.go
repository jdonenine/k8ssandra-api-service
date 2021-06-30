package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jdonenine/k8ssandra-api-service/pkg/auth"
	"k8s.io/client-go/kubernetes"
)

type ControllerManager struct {
	Kubeclient    *kubernetes.Clientset
	KubeNamespace string
	Router        *mux.Router
	AuthManager   *auth.AuthManager
}

func (manager *ControllerManager) Init() (handler http.Handler) {
	if manager.Kubeclient == nil {
		log.Println("Unable to initialize ControllerManager, Kubeclient is invalid.")
		return nil
	}
	if len(manager.KubeNamespace) < 1 {
		log.Println("Unable to initialize ControllerManager, KubeNamespace is invalid.")
		return nil
	}
	if manager.AuthManager == nil {
		log.Println("Unable to initialize ControllerManager, AuthManager is invalid.")
		return nil
	}
	router := mux.NewRouter()
	manager.Router = router
	manager.registerRoutes()
	return manager.Router
}

func (manager *ControllerManager) registerRoutes() {
	authController := AuthController{ExpiresInMinutes: 60}
	manager.Router.HandleFunc("/v1/auth/token", manager.AuthManager.WithAuthentication(http.HandlerFunc(authController.GetToken))).Methods("GET")

	cassandraDatacentersController := CassandraDatacentersController{Kubeclient: manager.Kubeclient, KubeNamespace: manager.KubeNamespace}
	manager.Router.HandleFunc("/v1/cassandra-datacenters", manager.AuthManager.WithAuthentication(http.HandlerFunc(cassandraDatacentersController.GetCassDcs))).Methods("GET")
	manager.Router.HandleFunc("/v1/cassandra-datacenters/{name}", manager.AuthManager.WithAuthentication(http.HandlerFunc(cassandraDatacentersController.GetCassDc))).Methods("GET")

	statefulSetsController := StatefulSetsController{Kubeclient: manager.Kubeclient, KubeNamespace: manager.KubeNamespace}
	manager.Router.HandleFunc("/v1/cassandra-datacenters/{name}/stateful-sets", manager.AuthManager.WithAuthentication(http.HandlerFunc(statefulSetsController.GetCassDcStatefulSets))).Methods("GET")

	manager.Router.NotFoundHandler = http.HandlerFunc(manager.notFound)
}

func (manager ControllerManager) notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
