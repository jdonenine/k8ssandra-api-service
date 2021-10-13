// Copyright DataStax, Inc.
// Please see the included license file for details.

package models

import (
	"encoding/json"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// ClusterLabel is the operator's label for the cluster name
	ClusterLabel = "cassandra.datastax.com/cluster"

	// DatacenterLabel is the operator's label for the datacenter name
	DatacenterLabel = "cassandra.datastax.com/datacenter"

	// SeedNodeLabel is the operator's label for the seed node state
	SeedNodeLabel = "cassandra.datastax.com/seed-node"

	// RackLabel is the operator's label for the rack name
	RackLabel = "cassandra.datastax.com/rack"

	// RackLabel is the operator's label for the rack name
	CassOperatorProgressLabel = "cassandra.datastax.com/operator-progress"

	// PromMetricsLabel is a service label that can be selected for prometheus metrics scraping
	PromMetricsLabel = "cassandra.datastax.com/prom-metrics"

	// DatacenterAnnotation is the operator's annotation for the datacenter name
	DatacenterAnnotation = DatacenterLabel

	// ConfigHashAnnotation is the operator's annotation for the hash of the ConfigSecret
	ConfigHashAnnotation = "cassandra.datastax.com/config-hash"

	// CassNodeState
	CassNodeState = "cassandra.datastax.com/node-state"

	// Progress states for status
	ProgressUpdating ProgressState = "Updating"
	ProgressReady    ProgressState = "Ready"

	// Default port numbers
	DefaultNativePort    = 9042
	DefaultInternodePort = 7000
)

// This type exists so there's no chance of pushing random strings to our progress status
type ProgressState string

type CassandraUser struct {
	SecretName string `json:"secretName"`
	Superuser  bool   `json:"superuser"`
}

// CassandraDatacenterSpec defines the desired state of a CassandraDatacenter
// +k8s:openapi-gen=true
type CassandraDatacenterSpec struct {
	// Important: Run "mage operator:sdkGenerate" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags:
	// https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Desired number of Cassandra server nodes
	// +kubebuilder:validation:Minimum=1
	Size int32 `json:"size"`

	// Version string for config builder,
	// used to generate Cassandra server configuration
	// +kubebuilder:validation:Pattern=(6\.8\.\d+)|(3\.11\.\d+)|(4\.0\.\d+)
	ServerVersion string `json:"serverVersion"`

	// Cassandra server image name.
	// More info: https://kubernetes.io/docs/concepts/containers/images
	ServerImage string `json:"serverImage,omitempty"`

	// Server type: "cassandra" or "dse"
	// +kubebuilder:validation:Enum=cassandra;dse
	ServerType string `json:"serverType"`

	// Does the Server Docker image run as the Cassandra user?
	DockerImageRunsAsCassandra *bool `json:"dockerImageRunsAsCassandra,omitempty"`

	// Config for the server, in YAML format
	// +kubebuilder:pruning:PreserveUnknownFields
	Config json.RawMessage `json:"config,omitempty"`

	// ConfigSecret is the name of a secret that contains configuration for Cassandra. The
	// secret is expected to have a property named config whose value should be a JSON
	// formatted string that should look like this:
	//
	//    config: |-
	//      {
	//        "cassandra-yaml": {
	//          "read_request_timeout_in_ms": 10000
	//        },
	//        "jmv-options": {
	//          "max_heap_size": 1024M
	//        }
	//      }
	//
	// ConfigSecret is mutually exclusive with Config. ConfigSecret takes precedence and
	// will be used exclusively if both properties are set. The operator sets a watch such
	// that an update to the secret will trigger an update of the StatefulSets.
	ConfigSecret string `json:"configSecret,omitempty"`

	// Config for the Management API certificates
	ManagementApiAuth ManagementApiAuthConfig `json:"managementApiAuth,omitempty"`

	//NodeAffinityLabels to pin the Datacenter, using node affinity
	NodeAffinityLabels map[string]string `json:"nodeAffinityLabels,omitempty"`

	// Kubernetes resource requests and limits, per pod
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`

	// Kubernetes resource requests and limits per system logger container.
	SystemLoggerResources corev1.ResourceRequirements `json:"systemLoggerResources,omitempty"`

	// Kubernetes resource requests and limits per server config initialization container.
	ConfigBuilderResources corev1.ResourceRequirements `json:"configBuilderResources,omitempty"`

	// A list of the named racks in the datacenter, representing independent failure domains. The
	// number of racks should match the replication factor in the keyspaces you plan to create, and
	// the number of racks cannot easily be changed once a datacenter is deployed.
	Racks []Rack `json:"racks,omitempty"`

	// Describes the persistent storage request of each server node
	StorageConfig StorageConfig `json:"storageConfig"`

	// A list of pod names that need to be replaced.
	ReplaceNodes []string `json:"replaceNodes,omitempty"`

	// The name by which CQL clients and instances will know the cluster. If the same
	// cluster name is shared by multiple Datacenters in the same Kubernetes namespace,
	// they will join together in a multi-datacenter cluster.
	// +kubebuilder:validation:MinLength=2
	ClusterName string `json:"clusterName"`

	// A stopped CassandraDatacenter will have no running server pods, like using "stop" with
	// traditional System V init scripts. Other Kubernetes resources will be left intact, and volumes
	// will re-attach when the CassandraDatacenter workload is resumed.
	Stopped bool `json:"stopped,omitempty"`

	// Container image for the config builder init container.
	ConfigBuilderImage string `json:"configBuilderImage,omitempty"`

	// Indicates that configuration and container image changes should only be pushed to
	// the first rack of the datacenter
	CanaryUpgrade bool `json:"canaryUpgrade,omitempty"`

	// The number of nodes that will be updated when CanaryUpgrade is true. Note that the value is
	// either 0 or greater than the rack size, then all nodes in the rack will get updated.
	CanaryUpgradeCount int32 `json:"canaryUpgradeCount,omitempty"`

	// Turning this option on allows multiple server pods to be created on a k8s worker node.
	// By default the operator creates just one server pod per k8s worker node using k8s
	// podAntiAffinity and requiredDuringSchedulingIgnoredDuringExecution.
	AllowMultipleNodesPerWorker bool `json:"allowMultipleNodesPerWorker,omitempty"`

	// This secret defines the username and password for the Cassandra server superuser.
	// If it is omitted, we will generate a secret instead.
	SuperuserSecretName string `json:"superuserSecretName,omitempty"`

	// The k8s service account to use for the server pods
	ServiceAccount string `json:"serviceAccount,omitempty"`

	// Whether to do a rolling restart at the next opportunity. The operator will set this back
	// to false once the restart is in progress.
	RollingRestartRequested bool `json:"rollingRestartRequested,omitempty"`

	// A map of label keys and values to restrict Cassandra node scheduling to k8s workers
	// with matchiing labels.
	// More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#nodeselector
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// Rack names in this list are set to the latest StatefulSet configuration
	// even if Cassandra nodes are down. Use this to recover from an upgrade that couldn't
	// roll out.
	ForceUpgradeRacks []string `json:"forceUpgradeRacks,omitempty"`

	DseWorkloads *DseWorkloads `json:"dseWorkloads,omitempty"`

	// PodTemplate provides customisation options (labels, annotations, affinity rules, resource requests, and so on) for the cassandra pods
	PodTemplateSpec *corev1.PodTemplateSpec `json:"podTemplateSpec,omitempty"`

	// Cassandra users to bootstrap
	Users []CassandraUser `json:"users,omitempty"`

	Networking *NetworkingConfig `json:"networking,omitempty"`

	AdditionalSeeds []string `json:"additionalSeeds,omitempty"`

	// Deprecated: Reaper's sidecar mode has too many problems in Kubernetes for it to
	// usable. In order for it to work reliably, changes in Reaper would be needed. See
	// https://github.com/thelastpickle/cassandra-reaper/issues/956 for details. Because
	// those changes were not implemented in Reaper and because Reaper support was instead
	// added through k8ssandra, this field will be removed in the 1.8.0 release.
	Reaper *ReaperConfig `json:"reaper,omitempty"`

	// Configuration for disabling the simple log tailing sidecar container. Our default is to have it enabled.
	DisableSystemLoggerSidecar bool `json:"disableSystemLoggerSidecar,omitempty"`

	// Container image for the log tailing sidecar container.
	SystemLoggerImage string `json:"systemLoggerImage,omitempty"`

	// AdditionalServiceConfig allows to define additional parameters that are included in the created Services. Note, user can override values set by cass-operator and doing so could break cass-operator functionality.
	// Avoid label "cass-operator" and anything that starts with "cassandra.datastax.com/"
	AdditionalServiceConfig ServiceConfig `json:"additionalServiceConfig,omitempty"`

	// Tolerations applied to the Cassandra pod. Note that these cannot be overridden with PodTemplateSpec.
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`
}

type NetworkingConfig struct {
	NodePort    *NodePortConfig `json:"nodePort,omitempty"`
	HostNetwork bool            `json:"hostNetwork,omitempty"`
}

type NodePortConfig struct {
	Native       int `json:"native,omitempty"`
	NativeSSL    int `json:"nativeSSL,omitempty"`
	Internode    int `json:"internode,omitempty"`
	InternodeSSL int `json:"internodeSSL,omitempty"`
}

type DseWorkloads struct {
	AnalyticsEnabled bool `json:"analyticsEnabled,omitempty"`
	GraphEnabled     bool `json:"graphEnabled,omitempty"`
	SearchEnabled    bool `json:"searchEnabled,omitempty"`
}

// StorageConfig defines additional storage configurations
type AdditionalVolumes struct {
	// Mount path into cassandra container
	MountPath string `json:"mountPath"`
	// Name of the pvc
	// +kubebuilder:validation:Pattern=[a-z0-9]([-a-z0-9]*[a-z0-9])?
	Name string `json:"name"`
	// Persistent volume claim spec
	PVCSpec corev1.PersistentVolumeClaimSpec `json:"pvcSpec"`
}

type AdditionalVolumesSlice []AdditionalVolumes

type StorageConfig struct {
	CassandraDataVolumeClaimSpec *corev1.PersistentVolumeClaimSpec `json:"cassandraDataVolumeClaimSpec,omitempty"`
	AdditionalVolumes            AdditionalVolumesSlice            `json:"additionalVolumes,omitempty"`
}

// ServiceConfig defines additional service configurations.
type ServiceConfig struct {
	DatacenterService     ServiceConfigAdditions `json:"dcService,omitempty"`
	SeedService           ServiceConfigAdditions `json:"seedService,omitempty"`
	AllPodsService        ServiceConfigAdditions `json:"allpodsService,omitempty"`
	AdditionalSeedService ServiceConfigAdditions `json:"additionalSeedService,omitempty"`
	NodePortService       ServiceConfigAdditions `json:"nodePortService,omitempty"`
}

// ServiceConfigAdditions exposes additional options for each service
type ServiceConfigAdditions struct {
	Labels      map[string]string `json:"additionalLabels,omitempty"`
	Annotations map[string]string `json:"additionalAnnotations,omitempty"`
}

// Rack ...
type Rack struct {
	// The rack name
	// +kubebuilder:validation:MinLength=2
	Name string `json:"name"`

	// Deprecated. Use nodeAffinityLabels instead. Zone name to pin the rack, using node affinity
	Zone string `json:"zone,omitempty"`

	//NodeAffinityLabels to pin the rack, using node affinity
	NodeAffinityLabels map[string]string `json:"nodeAffinityLabels,omitempty"`
}

type CassandraNodeStatus struct {
	HostID string `json:"hostID,omitempty"`
}

type CassandraStatusMap map[string]CassandraNodeStatus

type DatacenterConditionType string

const (
	DatacenterReady          DatacenterConditionType = "Ready"
	DatacenterInitialized    DatacenterConditionType = "Initialized"
	DatacenterReplacingNodes DatacenterConditionType = "ReplacingNodes"
	DatacenterScalingUp      DatacenterConditionType = "ScalingUp"
	DatacenterScalingDown    DatacenterConditionType = "ScalingDown"
	DatacenterUpdating       DatacenterConditionType = "Updating"
	DatacenterStopped        DatacenterConditionType = "Stopped"
	DatacenterResuming       DatacenterConditionType = "Resuming"
	DatacenterRollingRestart DatacenterConditionType = "RollingRestart"
	DatacenterValid          DatacenterConditionType = "Valid"
)

type DatacenterCondition struct {
	Type               DatacenterConditionType `json:"type"`
	Status             corev1.ConditionStatus  `json:"status"`
	Reason             string                  `json:"reason"`
	Message            string                  `json:"message"`
	LastTransitionTime metav1.Time             `json:"lastTransitionTime,omitempty"`
}

// CassandraDatacenterStatus defines the observed state of CassandraDatacenter
// +k8s:openapi-gen=true
type CassandraDatacenterStatus struct {
	Conditions []DatacenterCondition `json:"conditions,omitempty"`

	// Deprecated. Use usersUpserted instead. The timestamp at
	// which CQL superuser credentials were last upserted to the
	// management API
	// +optional
	SuperUserUpserted metav1.Time `json:"superUserUpserted,omitempty"`

	// The timestamp at which managed cassandra users' credentials
	// were last upserted to the management API
	// +optional
	UsersUpserted metav1.Time `json:"usersUpserted,omitempty"`

	// The timestamp when the operator last started a Server node
	// with the management API
	// +optional
	LastServerNodeStarted metav1.Time `json:"lastServerNodeStarted,omitempty"`

	// Last known progress state of the Cassandra Operator
	// +optional
	CassandraOperatorProgress ProgressState `json:"cassandraOperatorProgress,omitempty"`

	// +optional
	LastRollingRestart metav1.Time `json:"lastRollingRestart,omitempty"`

	// +optional
	NodeStatuses CassandraStatusMap `json:"nodeStatuses"`

	// +optional
	NodeReplacements []string `json:"nodeReplacements"`

	// +optional
	QuietPeriod metav1.Time `json:"quietPeriod,omitempty"`

	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CassandraDatacenter is the Schema for the cassandradatacenters API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=cassandradatacenters,scope=Namespaced,shortName=cassdc;cassdcs
type CassandraDatacenter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CassandraDatacenterSpec   `json:"spec,omitempty"`
	Status CassandraDatacenterStatus `json:"status,omitempty"`
}

type ManagementApiAuthManualConfig struct {
	ClientSecretName string `json:"clientSecretName"`
	ServerSecretName string `json:"serverSecretName"`
	// +optional
	SkipSecretValidation bool `json:"skipSecretValidation,omitempty"`
}

type ManagementApiAuthInsecureConfig struct {
}

type ManagementApiAuthConfig struct {
	Insecure *ManagementApiAuthInsecureConfig `json:"insecure,omitempty"`
	Manual   *ManagementApiAuthManualConfig   `json:"manual,omitempty"`
	// other strategy configs (e.g. Cert Manager) go here
}

type ReaperConfig struct {
	Enabled bool `json:"enabled,omitempty"`

	Image string `json:"image,omitempty"`

	ImagePullPolicy corev1.PullPolicy `json:"imagePullPolicy,omitempty"`

	// Kubernetes resource requests and limits per reaper container.
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CassandraDatacenterList contains a list of CassandraDatacenter
type CassandraDatacenterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CassandraDatacenter `json:"items"`
}
