/*
 * K8ssandra Cluster API
 *
 * A RESTful service providing control and visibility into a K8ssandra cluster.
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

type CassandraDatacenterSpec struct {

	ClusterName string `json:"clusterName,omitempty"`

	Size int32 `json:"size,omitempty"`

	ServerImage string `json:"serverImage,omitempty"`

	ServerType string `json:"serverType,omitempty"`

	ServerVersion string `json:"serverVersion,omitempty"`

	DockerImageRunsAsCassandra bool `json:"dockerImageRunsAsCassandra,omitempty"`
}