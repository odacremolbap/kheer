package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OperationSpec is the specification for a kheer operation object
type OperationSpec struct {
	Tasks []string `json:"tasks"`
}

// OperationStatus contain last known status for the operation
type OperationStatus struct {
	Status string `json:"status"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Operation is the CRD definition for the kheer operation object
type Operation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   OperationSpec   `json:"spec"`
	Status OperationStatus `json:"status"`
}
