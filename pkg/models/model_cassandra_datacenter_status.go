/*
 * K8ssandra Cluster API
 *
 * A RESTful service providing control and visibility into a K8ssandra cluster.
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models
import (
	"time"
)

type CassandraDatacenterStatus struct {

	CassandraOperatorProgress string `json:"cassandraOperatorProgress,omitempty"`

	LastServerNodeStarted time.Time `json:"lastServerNodeStarted,omitempty"`

	Conditions []StatusCondition `json:"conditions,omitempty"`
}
