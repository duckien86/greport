package common

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func ValidateStruct(data interface{}) (interface{}, error) {

	validate = validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("validPhone", cValidatePhoneNum) // regist custom phonenumber validation
	err := validate.Struct(data)
	details := make(map[string]string)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := strings.ToLower(err.Field())
			if len(err.Param()) > 0 {
				details[fieldName] = fmt.Sprintf("%s must be %s (%s)",
					fieldName, err.Tag(), err.Param())
			} else {
				details[fieldName] = fmt.Sprintf("%s must be %s %s",
					fieldName, err.Tag(), err.Param())
			}
		}
	}
	return details, err
}

// Custom func phonenumber validation
func cValidatePhoneNum(fl validator.FieldLevel) bool {
	pattern := `^0\d{9,10}$`
	// Compile the regex pattern
	regex := regexp.MustCompile(pattern)
	phoneNumber := fl.Field().String()
	// Match the phone number against the regex pattern
	return regex.MatchString(phoneNumber)
}
