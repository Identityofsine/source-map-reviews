package providers

import (
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/types"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

type GoogleAuthProvider struct {
}

// move this to validators?
func (obj *GoogleAuthProvider) validate(args AuthenticatorArgs) bool {
	//check if args is nil or what we expect from a google auth request
	if args == nil {
		return false
	}
	return true
}

func (obj *GoogleAuthProvider) Authenticate(args AuthenticatorArgs) (*Token, db.DatabaseError) {

	if valid := obj.validate(args); !valid {
		return nil, db.NewDatabaseError("GoogleAuthProvider::Authenticate", "args is nil", "args-nil", 400)
	}

	return nil, db.NewDatabaseError("GoogleAuthProvider::Authenticate", "not implemented", "not-implemented", 501)
}

func (obj *GoogleAuthProvider) Name() string {
	return "Google"
}
