package providers

import . "github.com/identityofsine/fofx-go-gin-api-template/pkg/register/types"
import . "github.com/identityofsine/fofx-go-gin-api-template/pkg/register/validator"

type internalRegisterValidator struct {
}

const source = "InternalRegisterValidator"

func (obj *internalRegisterValidator) Validate(args RegisterArgs) ValidatorError {
	//check if args is nil or what we expect from a register request
	if args == nil {
		return NewValidationError(source, "args is nil", 400)
	}
	if args["username"] == "" {
		return NewValidationError(source, "email is empty", 400)
	}
	if args["password"] == "" {
		return NewValidationError(source, "password is empty", 400)
	}

	return nil
}

var (
	InternalRegisterValidator = &internalRegisterValidator{}
)
