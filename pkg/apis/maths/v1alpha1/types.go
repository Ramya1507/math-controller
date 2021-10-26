package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)




// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MathResource is a specification for a custom resource
type MathResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MathResourceSpec   `json:"spec"`
  Status MathResourceStatus `json:"status"`
}

// MathResourceSpec is the spec for a custom resource
type MathResourceSpec struct {
	FirstNum   int32        `json:"firstNum"`
	SecondNum  int32       `json:"secondNum"`
	Operation  string      `json:"operation"`

}

type MathResourceStatus struct {
	State   string `json:"state"`
	Message string `json:"message"`
}



// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MathResourceList is a list of custom resources
type MathResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []MathResource `json:"items"`
}
