package providers

import (
	"errors"

	userdb "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
	authTypeLks "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/types"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/register/types"
	validators "github.com/identityofsine/fofx-go-gin-api-template/pkg/register/validator/providers"
)

type InternalRegisterProvider struct {
}

func (obj *InternalRegisterProvider) Register(args RegisterArgs) error {
	err := validators.InternalRegisterValidator.Validate(args)
	if err != nil {
		return err
	}

	//create a new user if the username doesn't exist already
	if usr, err := userdb.GetUserByUsername(args["username"].(string)); err != nil || usr != nil {
		if err != nil && err.Code != 404 {
			return err
		} else if err != nil && err.Code == 404 {
			//nop
		} else {
			return errors.New("user already exists")
		}
	}

	return userdb.CreateUser(args["username"].(string), args["password"].(string), authTypeLks.AUTH_METHOD_INTERNAL)
}

func (obj *InternalRegisterProvider) Name() string {
	return "Internal"
}
