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

type StatusCondition struct {

	LastTransitionTime time.Time `json:"lastTransitionTime,omitempty"`

	Message string `json:"message,omitempty"`

	Reason string `json:"reason,omitempty"`

	Status string `json:"status,omitempty"`

	Type_ string `json:"type,omitempty"`
}