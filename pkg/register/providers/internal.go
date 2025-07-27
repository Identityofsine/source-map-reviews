package providers

import (
	"errors"

	userdb "github.com/identityofsine/fofx-go-gin-api-template/internal/repository"
	authTypeLks "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/authtypes"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/bcrypt"
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

	password, ok := args["password"].(string)
	if !ok || password == "" {
		return errors.New("password is required")
	}

	password, bcryptError := bcrypt.HashString(password)
	if bcryptError != nil {
		return bcryptError
	}

	derr := userdb.CreateUser(args["username"].(string), password, authTypeLks.AUTH_METHOD_INTERNAL)
	if derr != nil {
		return derr
	}

	return nil
}

func (obj *InternalRegisterProvider) Name() string {
	return "Internal"
}
