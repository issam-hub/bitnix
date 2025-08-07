package validator

type ValidatorError struct {
	Field   string
	Message string
}

type ValidationErrors []ValidatorError

func (ve ValidationErrors) HasErrors() bool {
	return len(ve) > 0
}
