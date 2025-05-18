package providers

import . "github.com/identityofsine/fofx-go-gin-api-template/pkg/register/types"
import validators "github.com/identityofsine/fofx-go-gin-api-template/pkg/register/validator/providers"

type InternalRegisterProvider struct {
}

func (obj *InternalRegisterProvider) Register(args RegisterArgs) error {
	err := validators.InternalRegisterValidator.Validate(args)
	return err
}

func (obj *InternalRegisterProvider) Name() string {
	return "Internal"
}
