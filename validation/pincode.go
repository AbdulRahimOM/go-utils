package validation

import (
	"github.com/go-playground/validator/v10"
)

func init() {
	validate.RegisterValidation("pincode", validatePincode)
}

func validatePincode(fl validator.FieldLevel) bool {
	value := fl.Field().Uint()

	return (value >= 110000 && value <= 899999)
}
