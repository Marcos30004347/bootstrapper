package v1beta1

import "k8s.io/apimachinery/pkg/runtime"

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}

func SetDefaults_FooSpec(obj *FooSpec) {
	if len(obj.Bar) == 0 {
		obj.Bar = []FooBar{
			{"foo0", 1},
			{"foo1", 1},
			{"foo2", 1},
		}
	}

	for i := range obj.Bar {
		if obj.Bar[i].Quantity == 0 {
			obj.Bar[i].Quantity = 1
		}
	}
}
