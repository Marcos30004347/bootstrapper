package validation

import (
	"github.com/marcos30004347/kubeapi/pkg/apis/restaurant"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func ValidateFoo(f *restaurant.Foo) field.ErrorList {
	allErrs := field.ErrorList{}

	allErrs = append(allErrs, ValidateFooSpec(&f.Spec, field.NewPath("spec"))...)

	return allErrs
}

func ValidateFooSpec(s *restaurant.FooSpec, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	prevNames := map[string]bool{}
	for i := range s.Bar {
		if s.Bar[i].Quantity <= 0 {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("bars").Index(i).Child("quantity"), s.Bar[i].Quantity, "cannot be negative or zero"))
		}
		if len(s.Bar[i].Name) == 0 {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("bars").Index(i).Child("name"), s.Bar[i].Name, "cannot be empty"))
		} else {
			if prevNames[s.Bar[i].Name] {
				allErrs = append(allErrs, field.Invalid(fldPath.Child("bars").Index(i).Child("name"), s.Bar[i].Name, "must be unique"))
			}
			prevNames[s.Bar[i].Name] = true
		}
	}

	return allErrs
}
