package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	localauth "github.com/jdonenine/k8ssandra-api-service/pkg/auth"
	"github.com/jdonenine/k8ssandra-api-service/pkg/controllers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeNamespace string
var kubeconfigPath string
var kubeServerMasterUrl string

var serverPort string
var serverBindAddress string

var userSecretName string

var kubeclient *kubernetes.Clientset

func main() {
	log.Println("Starting k8ssandra-api-service...")
	initEnvironment()
	initKubeClient()
	authManager := initAuth()
	routeHandler := initContollers(&authManager)
	startServer(routeHandler)
}

func initEnvironment() {
	kubeconfigPath = os.Getenv("K8SSANDRA_API_SERVICE_KUBECONFIG_PATH")
	log.Printf("Initializing with K8SSANDRA_API_SERVICE_KUBECONFIG_PATH=%s", kubeconfigPath)

	kubeServerMasterUrl = os.Getenv("K8SSANDRA_API_SERVICE_KUBE_SERVER_MASTER_URL")
	log.Printf("Initializing with K8SSANDRA_API_SERVICE_KUBE_SERVER_MASTER_URL=%s", kubeServerMasterUrl)

	serverPort = os.Getenv("K8SSANDRA_API_SERVICE_PORT")
	log.Printf("Initializing with K8SSANDRA_API_SERVICE_PORT=%s", serverPort)
	if len(serverPort) < 1 {
		log.Printf("No K8SSANDRA_API_SERVICE_PORT set, starting with default '%s'", "3000")
		serverPort = "3000"
	}

	serverBindAddress = os.Getenv("K8SSANDRA_API_SERIVCE_BIND_ADDRESS")
	log.Printf("Initializing with K8SSANDRA_API_SERIVCE_BIND_ADDRESS=%s", serverBindAddress)
	if len(serverBindAddress) < 1 {
		log.Printf("No K8SSANDRA_API_SERIVCE_BIND_ADDRESS set, starting with default '%s'", "0.0.0.0")
		serverBindAddress = "0.0.0.0"
	}

	userSecretName = os.Getenv("K8SSANDRA_API_SERIVCE_USER_SERCRET_NAME")
	log.Printf("Initializing with K8SSANDRA_API_SERIVCE_USER_SERCRET_NAME=%s", userSecretName)
	if len(userSecretName) < 1 {
		log.Printf("No K8SSANDRA_API_SERIVCE_USER_SERCRET_NAME set, starting with default '%s'", "k8ssandra-api-service-user")
		userSecretName = "k8ssandra-api-service-user"
	}

	kubeNamespace = os.Getenv("K8SSANDRA_API_SERIVCE_NAMESPACE")
	log.Printf("Initializing with K8SSANDRA_API_SERIVCE_NAMESPACE=%s", kubeNamespace)
	if len(kubeNamespace) < 1 {
		log.Printf("No K8SSANDRA_API_SERIVCE_NAMESPACE set, starting with default '%s'", "default")
		kubeNamespace = "default"
	}

}

func initKubeClient() {
	var clientConfig *rest.Config
	if len(kubeconfigPath) > 0 {
		// creates an out-of-cluster config
		config, err := clientcmd.BuildConfigFromFlags(kubeServerMasterUrl, kubeconfigPath)
		if err != nil {
			log.Fatalf("Unable to build client configuration from kubeconfig at '%s', failed with error: '%s'", kubeconfigPath, err)
		} else {
			clientConfig = config
		}
	} else {
		// creates the in-cluster config
		config, err := rest.InClusterConfig()
		if err != nil {
			log.Fatalf("Unable to build client configuration in-cluster: '%s'", err)
		} else {
			clientConfig = config
		}
	}
	if clientConfig == nil {
		log.Fatalln("Unable to build client configuration")
	}
	clientset, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		log.Fatalf("Unable to build client, failed with error: '%s'", err)
	} else {
		kubeclient = clientset
	}
}

func initAuth() (authManager localauth.AuthManager) {
	manager := localauth.AuthManager{UserSecretName: userSecretName, Kubeclient: kubeclient, KubeNamespace: kubeNamespace}
	manager.Init()
	return manager
}

func initContollers(authManager *localauth.AuthManager) (handler http.Handler) {
	if authManager == nil {
		log.Fatalln("Unable to initialize controllers, authManager is invalid.")
	}
	manager := controllers.ControllerManager{Kubeclient: kubeclient, KubeNamespace: kubeNamespace, AuthManager: authManager}
	return manager.Init()
}

func startServer(routeHandler http.Handler) {
	if routeHandler == nil {
		log.Fatalln("Unable to start server, routeHandler is invalid.")
	}
	serverBindAddressPort := strings.TrimSpace(serverBindAddress) + ":" + strings.TrimSpace(serverPort)
	log.Printf("Starting server listenting on http://%s", serverBindAddressPort)
	http.ListenAndServe(serverBindAddressPort, routeHandler)
}
