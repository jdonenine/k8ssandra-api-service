package models

import "time"

type CassandraDatacenterStatus struct {
	CassandraOperatorProgress string               `json:"cassandraOperatorProgress"`
	LastServerNodeStarted     time.Time            `json:"lastServerNodeStarted"`
	Conditions                []K8sStatusCondition `json:"conditions"`
}

type CassandraDatacenter struct {
	ApiVersion string                    `json:"apiVersion"`
	Kind       string                    `json:"kind"`
	Metadata   K8sMetadata               `json:"metadata"`
	Status     CassandraDatacenterStatus `json:"status"`
}

type CassandraDatacenters struct {
	ApiVersion string                `json:"apiVersion"`
	Items      []CassandraDatacenter `json:"items"`
	Kind       string                `json:"kind"`
}
