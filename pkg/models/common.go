package models

import "time"

type ResponseMetadata struct {
	Count int `json:"count"`
	Total int `json:"total"`
}

type K8sMetadata struct {
	Name        string            `json:"name"`
	Namespace   string            `json:"namespace"`
	Annotations map[string]string `json:"annotations"`
	Labels      map[string]string `json:"labels"`
	Generation  int               `json:"generation"`
}

type K8sStatusCondition struct {
	LastTransitionTime time.Time `json:"lastTransitionTime"`
	Message            string    `json:"message"`
	Reason             string    `json:"reason"`
	Status             string    `json:"status"`
	Type               string    `json:"type"`
}
