package providers

import (
	userDto "github.com/identityofsine/fofx-go-gin-api-template/api/dto/user"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/components/user/model"
	userService "github.com/identityofsine/fofx-go-gin-api-template/internal/components/user/service"
	userdb "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/model"
	tokenService "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/service"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/types"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
)

//Internal Auth works as a system to authenticate users directly with the system.
//This is the classic username and password authentication system.

type InternalAuthProvider struct {
}

// TODO: change bool to include error reason
// move this to validators?
func (obj *InternalAuthProvider) validate(args AuthenticatorArgs) bool {
	//check if args is nil or what we expect from a google auth request
	if args == nil {
		return false
	}
	if args.Keys["username"] == nil || args.Keys["password"] == nil {
		return false
	}

	return true
}

func (obj *InternalAuthProvider) Authenticate(args AuthenticatorArgs) (*Token, db.DatabaseError) {
	if valid := obj.validate(args); !valid {
		return nil, db.NewDatabaseError("InternalAuthProvider::Authenticate", "args is nil", "args-nil", 400)
	}

	userdb, derr := userdb.GetUserByUsername(args.Keys["username"].(string))
	if derr != nil || userdb == nil {
		return nil, derr
	}

	user := userDto.Map(*userdb)

	if userService.IsPasswordsEqual(user, User{
		Username: args.Keys["username"].(string),
		Password: args.Keys["password"].(string),
	}) == false {
		return nil, db.NewDatabaseError("error comparing passwords", "passwords do not match", "passwords-do-not-match", 401)
	}

	token, err := tokenService.CreateLoginToken(user.ID)
	if err != nil {
		storedlogs.LogError("error creating token", err)
		return nil, db.NewDatabaseError("error creating token", "error creating token", "error-creating-token", 500)
	}

	return token, nil
}

func (obj *InternalAuthProvider) Name() string {
	return "Internal"
}
