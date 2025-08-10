package validator

import "strings"

type ValidationError struct {
	Field   string
	Message string
}

type ValidationErrors []ValidationError

func (ve ValidationErrors) HasErrors() bool {
	return len(ve) > 0
}

func (ve ValidationErrors) Error() string {
	var messages []string
	for _, err := range ve {
		messages = append(messages, err.Field+": "+err.Message)
	}
	return strings.Join(messages, ", ")
}

func (ve ValidationErrors) ToMap() map[string]string {
	errors := make(map[string]string)

	for _, err := range ve {
		if _, ok := errors[err.Field]; ok {
			errors[err.Field] = ", " + err.Message
		} else {
			errors[err.Field] = err.Message
		}
	}
	return errors
}
