package resources

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/jdonenine/k8ssandra-api-service/pkg/models"
	"k8s.io/client-go/kubernetes"
)

type ResourceManager struct {
	Kubeclient    *kubernetes.Clientset
	KubeNamespace string
}

func (manager *ResourceManager) GetCassDcs() (*models.CassandraDatacenterList, error) {
	if manager.Kubeclient == nil {
		return nil, errors.New("the Kubeclient provided is not valid")
	}

	// Query the k8s API for all of the CassandraDataCenter resources in the given namespace
	resourcePath := fmt.Sprintf("/apis/cassandra.datastax.com/v1beta1/namespaces/%s/cassandradatacenters/", manager.KubeNamespace)
	rawCassDcs, getErr := manager.Kubeclient.RESTClient().
		Get().
		AbsPath(resourcePath).
		DoRaw(context.TODO())
	if getErr != nil {
		log.Printf("Unable to retrieve CassandraDatacenterList resource from namesapce '%s', failed with error: '%s'", manager.KubeNamespace, getErr)
		return nil, getErr
	}

	cassDcs := models.CassandraDatacenterList{}
	unmarshalErr := json.Unmarshal([]byte(rawCassDcs), &cassDcs)
	if unmarshalErr != nil {
		log.Printf("Unable to process CassandraDatacenterList resource from namesapce '%s', failed with error: '%s'", manager.KubeNamespace, unmarshalErr)
		return nil, unmarshalErr
	}

	return &cassDcs, nil
}

func (manager *ResourceManager) GetCassDc(name string) (*models.CassandraDatacenter, error) {
	if manager.Kubeclient == nil {
		return nil, errors.New("the Kubeclient provided is not valid")
	}

	// Query the k8s API for the CassandraDataCenter resource in the given namespace
	resourcePath := fmt.Sprintf("/apis/cassandra.datastax.com/v1beta1/namespaces/%s/cassandradatacenters/%s", manager.KubeNamespace, name)
	rawCassDc, getErr := manager.Kubeclient.RESTClient().
		Get().
		AbsPath(resourcePath).
		DoRaw(context.TODO())
	if getErr != nil {
		log.Printf("Unable to retrieve CassandraDatacenter resource with name '%s' from namesapce '%s', failed with error: '%s'", name, manager.KubeNamespace, getErr)
		return nil, getErr
	}

	cassDc := models.CassandraDatacenter{}
	unmarshalErr := json.Unmarshal([]byte(rawCassDc), &cassDc)
	if unmarshalErr != nil {
		log.Printf("Unable to process CassandraDatacenter resource with name '%s' from namesapce '%s', failed with error: '%s'", name, manager.KubeNamespace, unmarshalErr)
		return nil, unmarshalErr
	}

	return &cassDc, nil
}
