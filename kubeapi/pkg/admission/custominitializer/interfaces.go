package custominitializer

import (
	"k8s.io/apiserver/pkg/admission"

	informers "github.com/marcos30004347/kubeapi/pkg/generated/informers/externalversions"
)

// WantsRestaurantInformerFactory defines a function which sets InformerFactory for admission plugins that need it
type WantsRestaurantInformerFactory interface {
	SetRestaurantInformerFactory(informers.SharedInformerFactory)
	admission.InitializationValidator
}
