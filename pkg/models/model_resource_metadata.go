/*
 * K8ssandra Cluster API
 *
 * A RESTful service providing control and visibility into a K8ssandra cluster.
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

type ResourceMetadata struct {

	Name string `json:"name,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	Generation int32 `json:"generation,omitempty"`

	Annotations map[string]string `json:"annotations,omitempty"`

	Labels map[string]string `json:"labels,omitempty"`
}
