package common

import (
	"regexp"
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

func IsMobileNumber(mobile string) bool {
	if len(mobile) != 10 {
		return false
	}
	for _, char := range mobile {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func IsEmail(email string) bool {
	// Define a regular expression for validating an email address
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	// Use MatchString method to validate the email string
	return emailRegex.MatchString(email)
}

func ValidatePassword(fl validator.FieldLevel) bool {
	//
	return true
}
