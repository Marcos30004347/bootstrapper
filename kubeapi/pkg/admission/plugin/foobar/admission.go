package foobar

import (
	"context"
	"fmt"
	"io"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apiserver/pkg/admission"

	"github.com/marcos30004347/kubeapi/pkg/admission/custominitializer"
	"github.com/marcos30004347/kubeapi/pkg/apis/restaurant"
	informers "github.com/marcos30004347/kubeapi/pkg/generated/informers/externalversions"
	listers "github.com/marcos30004347/kubeapi/pkg/generated/listers/restaurant/v1alpha1"
)

// Register registers a plugin
func Register(plugins *admission.Plugins) {
	plugins.Register("FooBar", func(config io.Reader) (admission.Interface, error) {
		return New()
	})
}

type FooBarPlugin struct {
	*admission.Handler
	barLister listers.BarLister
}

var _ = custominitializer.WantsRestaurantInformerFactory(&FooBarPlugin{})

// Admit ensures that the object in-flight is of kind Foo.
// In addition checks that the toppings are known.
func (d *FooBarPlugin) Admit(ctx context.Context, a admission.Attributes, _ admission.ObjectInterfaces) error {
	// we are only interested in pizzas
	if a.GetKind().GroupKind() != restaurant.Kind("Foo") {
		return nil
	}

	if !d.WaitForReady() {
		return admission.NewForbidden(a, fmt.Errorf("not yet ready to handle request"))
	}

	obj := a.GetObject()
	foo := obj.(*restaurant.Foo)
	for _, top := range foo.Spec.Bar {
		if _, err := d.barLister.Get(top.Name); err != nil && errors.IsNotFound(err) {
			return admission.NewForbidden(
				a,
				fmt.Errorf("unknown topping: %s", top.Name),
			)
		}
	}

	return nil
}

// SetRestaurantInformerFactory gets Lister from SharedInformerFactory.
// The lister knows how to lists Bar.
func (d *FooBarPlugin) SetRestaurantInformerFactory(f informers.SharedInformerFactory) {
	d.barLister = f.Restaurant().V1alpha1().Bars().Lister()
	d.SetReadyFunc(f.Restaurant().V1alpha1().Bars().Informer().HasSynced)
}

// ValidaValidateInitializationte checks whether the plugin was correctly initialized.
func (d *FooBarPlugin) ValidateInitialization() error {
	if d.barLister == nil {
		return fmt.Errorf("missing policy lister")
	}
	return nil
}

// New creates a new ban foo topping admission plugin
func New() (*FooBarPlugin, error) {
	return &FooBarPlugin{
		Handler: admission.NewHandler(admission.Create, admission.Update),
	}, nil
}
