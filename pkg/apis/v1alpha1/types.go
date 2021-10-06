package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const GroupName = "mathcontroller"


// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CustomResource is a specification for a Foo resource
type CustomResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CustomResourceSpec   `json:"spec"`
	Status CustomResourceStatus `json:"status"`
}

// CustomResourceSpec is the spec for a Foo resource
type CustomResourceSpec struct {
	FirstNum int32        `json:"firstNum"`
	SecondNum int32       `json:"secondNum"`
	Operation string      `json:"operation"`

}

// CustomResourceStatus is the status for a Foo resource
type CustomResourceStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CustomResourceList is a list of Foo resources
type CustomResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []CustomResource `json:"items"`
}
