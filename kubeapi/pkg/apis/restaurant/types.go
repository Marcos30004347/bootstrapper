/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package restaurant

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Foo specifies an offered pizza with toppings.
type Foo struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   FooSpec
	Status FooStatus
}

type FooSpec struct {
	// +k8s:conversion-gen=false
	// toppings is a list of Bar names. They don't have to be unique. Order does not matter.
	Bar []FooBar
}

type FooBar struct {
	// name is the name of a Bar object .
	Name string
	// quantity is the number of how often the topping is put onto the pizza.
	Quantity int
}

type FooStatus struct {
	// cost is the cost of the whole pizza including all toppings.
	Cost float64
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooList is a list of Foo objects.
type FooList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []Foo
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Bar is a topping put onto a pizza.
type Bar struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec BarSpec
}

type BarSpec struct {
	// cost is the cost of one instance of this topping.
	Cost float64
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BarList is a list of Bar objects.
type BarList struct {
	metav1.TypeMeta
	metav1.ListMeta

	// Items is a list of Bars
	Items []Bar
}
