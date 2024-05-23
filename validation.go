package credmark

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var (
	messages = map[string]string{
		"required":         "is_required",
		"required_if":      "is_required",
		"required_without": "is_required",
		"email":            "invalid_email",
		"oneof":            "value_is_invalid",
		"min":              "less_than_min",
		"max":              "over_max",
		"e164":             "is_not_e164",
	}
)

// GetNiceMessage returns a formatted message based on namespace and tag
func getNiceMessage(namespace, tag string) string {
	return fmt.Sprintf("%s %s", namespace, messages[tag])
}

// ValidateStruct will validate the reqPayload based on the struct tags
func ValidateStruct(reqPayload interface{}) error {
	v := validator.New()
	err := v.Struct(reqPayload)

	if err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return errors.New("invalid JSON body")
		}

		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(getNiceMessage(err.Namespace(), err.Tag()))
		}
	}

	return nil
}
