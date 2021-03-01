package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Foo specifies an offered Foo with bar.
type Foo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   FooSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status FooStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type FooSpec struct {
	// +k8s:conversion-gen=false
	// bar is a list of Bar names. They don't have to be unique. Order does not matter.
	Bar []string `json:"bar" protobuf:"bytes,1,rep,name=bar"`
}

type FooStatus struct {
	// cost is the cost of the whole Foo including all bar.
	Cost float64 `json:"cost,omitempty" protobuf:"bytes,1,opt,name=cost"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooList is a list of Foo objects.
type FooList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []Foo `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Bar
type Bar struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec BarSpec
}

type BarSpec struct {
	// cost is the cost of one instance of this topping.
	Cost float64 `json:"cost" protobuf:"bytes,1,name=cost"`
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BarList is a list of Bar objects.
type BarList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []Bar `json:"items" protobuf:"bytes,2,rep,name=items"`
}
