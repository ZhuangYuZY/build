package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterBuildStrategy is the Schema for the clusterbuildstrategies API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=clusterbuildstrategies,scope=Cluster,shortName=cbs;cbss
type ClusterBuildStrategy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BuildStrategySpec   `json:"spec,omitempty"`
	Status BuildStrategyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterBuildStrategyList contains a list of ClusterBuildStrategy
type ClusterBuildStrategyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterBuildStrategy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterBuildStrategy{}, &ClusterBuildStrategyList{})
}
