package v1beta1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true

// Foo specifies an offered Foo with toppings.
type Foo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              FooSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            FooStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type FooSpec struct {
	// bar is a list of Bars names. They don't have to be unique. Order does not matter.

	// +listType=set
	Bar []FooBar `json:"bar" protobuf:"bytes,1,rep,name=bar"`
}

type FooBar struct {
	// name is the name of a Bar object .
	Name string `json:"name" protobuf:"bytes,1,name=name"`
	// quantity is the number of how often the topping is put onto the Foo.
	// +optional
	Quantity int `json:"quantity" protobuf:"bytes,2,opt,name=quantity"`
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
