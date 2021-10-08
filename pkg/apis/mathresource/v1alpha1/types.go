package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)




// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MathResource is a specification for a Foo resource
type MathResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MathResourceSpec   `json:"spec"`

}

// MathResourceSpec is the spec for a Foo resource
type MathResourceSpec struct {
	FirstNum   int32        `json:"firstNum"`
	SecondNum  int32       `json:"secondNum"`
	Operation  string      `json:"operation"`

}



// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MathResourceList is a list of Foo resources
type MathResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []MathResource `json:"items"`
}
