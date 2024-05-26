package common

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = initValidator()

func initValidator() *validator.Validate {
	v := validator.New(validator.WithRequiredStructEnabled())
	v.RegisterValidation("mobile", ValidateMobileNumber)
	v.RegisterValidation("password", ValidatePassword)
	return v
}

func ValidateMobileNumber(fl validator.FieldLevel) bool {

	if len(fl.Field().String()) != 10 {
		return false
	}
	for _, char := range fl.Field().String() {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func ValidatePassword(fl validator.FieldLevel) bool {
	//
	return true
}
