// type RESTCreateStrategy interface {
//     runtime.ObjectTyper
//     // The name generator is used when the standard GenerateName field is set.
//     // The NameGenerator will be invoked prior to validation.
//     names.NameGenerator

//     // NamespaceScoped returns true if the object must be within a namespace.
//     NamespaceScoped() bool
//     // PrepareForCreate is invoked on create before validation to normalize
//     // the object. For example: remove fields that are not to be persisted,
//     // sort order-insensitive list fields, etc. This should not remove fields
//     // whose presence would be considered a validation error.
//     //
//     // Often implemented as a type check and an initailization or clearing of
//     // status. Clear the status because status changes are internal. External
//     // callers of an api (users) should not be setting an initial status on
//     // newly created objects.
//     PrepareForCreate(ctx context.Context, obj runtime.Object)
//     // Validate returns an ErrorList with validation errors or nil. Validate
//     // is invoked after default fields in the object have been filled in
//     // before the object is persisted. This method should not mutate the
//     // object.
//     Validate(ctx context.Context, obj runtime.Object) field.ErrorList
//     // Canonicalize allows an object to be mutated into a canonical form. This
//     // ensures that code that operates on these objects can rely on the common
//     // form for things like comparison. Canonicalize is invoked after
//     // validation has succeeded but before the object has been persisted.
//     // This method may mutate the object. Often implemented as a type check or
//     // empty method.
//     Canonicalize(obj runtime.Object)
// }

package bar

import (
	"context"
	"fmt"

	"github.com/marcos30004347/kubeapi/pkg/apis/restaurant"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
)

// NewStrategy creates and returns a barStrategy instance
func NewStrategy(typer runtime.ObjectTyper) barStrategy {
	return barStrategy{typer, names.SimpleNameGenerator}
}

// GetAttrs returns labels.Set, fields.Set, the presence of Initializers if any
// and error in case the given runtime.Object is not a Bar
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	apiserver, ok := obj.(*restaurant.Bar)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not a Bar")
	}
	return labels.Set(apiserver.ObjectMeta.Labels), SelectableFields(apiserver), nil
}

// MatchBar is the filter used by the generic etcd backend to watch events
// from etcd to clients of the apiserver only interested in specific labels/fields.
func MatchBar(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// SelectableFields returns a field set that represents the object.
func SelectableFields(obj *restaurant.Bar) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
}

type barStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (barStrategy) NamespaceScoped() bool {
	return true
}

func (barStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

func (barStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
}

// Here is where we actually use the Validate Function defined in the api
func (barStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}

}

func (barStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (barStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (barStrategy) Canonicalize(obj runtime.Object) {
}

func (barStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}
