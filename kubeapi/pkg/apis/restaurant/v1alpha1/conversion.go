/*
Copyright 2018 The Kubernetes Authors.

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

package v1alpha1

import (
	"github.com/marcos30004347/kubeapi/pkg/apis/restaurant"
	"k8s.io/apimachinery/pkg/conversion"
)

// Convert_v1alpha1_FooSpec_To_restaurant_FooSpec is an autogenerated conversion function.
func Convert_v1alpha1_FooSpec_To_restaurant_FooSpec(in *FooSpec, out *restaurant.FooSpec, s conversion.Scope) error {
	idx := map[string]int{}
	for _, top := range in.Bar {
		if i, duplicate := idx[top]; duplicate {
			out.Bar[i].Quantity++
			continue
		}
		idx[top] = len(out.Bar)
		out.Bar = append(out.Bar, restaurant.FooBar{
			Name:     top,
			Quantity: 1,
		})
	}

	return nil
}

// Convert_restaurant_FooSpec_To_v1alpha1_FooSpec is an autogenerated conversion function.
func Convert_restaurant_FooSpec_To_v1alpha1_FooSpec(in *restaurant.FooSpec, out *FooSpec, s conversion.Scope) error {
	for i := range in.Bar {
		for j := 0; j < in.Bar[i].Quantity; j++ {
			out.Bar = append(out.Bar, in.Bar[i].Name)
		}
	}

	return nil
}
