package providers

import (
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/components/user"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	repository "github.com/identityofsine/fofx-go-gin-api-template/internal/repository"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/authtypes"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/model"
	tokenService "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/service"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
)

//Internal Auth works as a system to authenticate users directly with the system.
//This is the classic username and password authentication system.

type InternalAuthProvider struct {
}

// TODO: change bool to include error reason
// move this to validators?
func (obj *InternalAuthProvider) validate(args authtypes.AuthenticatorArgs) bool {
	//check if args is nil or what we expect from a google auth request
	if args == nil {
		return false

	}
	if args.Keys["username"] == nil || args.Keys["password"] == nil {
		return false
	}

	return true
}

func (obj *InternalAuthProvider) Authenticate(args authtypes.AuthenticatorArgs) (*Token, authtypes.AuthError) {
	if valid := obj.validate(args); !valid {
		return nil, authtypes.NewAuthError(
			"InternalAuthProvider::Authenticate",
			"invalid-credentials",
			"invalid-credentials",
			exception.CODE_BAD_REQUEST,
		)
	}

	userdb, derr := repository.GetUserByUsername(args.Keys["username"].(string))
	if derr != nil || userdb == nil {
		return nil, authtypes.NewAuthError(
			"InternalAuthProvider::Authenticate",
			"invalid-credentials",
			"invalid-credentials",
			exception.CODE_UNAUTHORIZED,
		)
	}

	user := dbmapper.MapDbFields[repository.UserDB, User](*userdb)

	if IsPasswordsEqual(*user, User{
		Username: args.Keys["username"].(string),
		Password: args.Keys["password"].(string),
	}) == false {
		return nil, authtypes.NewAuthError(
			"InternalAuthProvider::Authenticate",
			"invalid-credentials",
			"invalid-credentials",
			exception.CODE_UNAUTHORIZED,
		)
	}

	token, err := tokenService.CreateLoginToken(user.ID)
	if err != nil {
		storedlogs.LogError("error creating token", err)
		return nil, authtypes.NewAuthError(
			"InternalAuthProvider::Authenticate",
			"error-creating-token",
			"error-creating-token",
			exception.CODE_INTERNAL_SERVER_ERROR,
		)
	}

	return token, nil
}

func (obj *InternalAuthProvider) Name() string {
	return "Internal"
}
