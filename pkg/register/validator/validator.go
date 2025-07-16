package validator

import . "github.com/identityofsine/fofx-go-gin-api-template/pkg/register/types"

type validationError struct {
	ValidatorSource string
	Message         string
	Code            int
}

func NewValidationError(source string, message string, code int) ValidatorError {
	return &validationError{
		ValidatorSource: source,
		Message:         message,
		Code:            code,
	}
}

func (e validationError) Error() string {
	return e.Message
}

type ValidatorError = *validationError

type Validator interface {
	Validate(RegisterArgs) ValidatorError
}
