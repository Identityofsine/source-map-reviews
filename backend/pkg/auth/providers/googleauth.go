package providers

import (
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/authtypes"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/authtypes"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/model"
	GoogleAuth "github.com/identityofsine/fofx-go-gin-api-template/pkg/oauth/providers/google"
)

type GoogleAuthProvider struct {
}

// TODO: replace db.DatabaseError with a more specific error type, with actionable error messages
func (obj *GoogleAuthProvider) Authenticate(args AuthenticatorArgs) (*Token, authtypes.AuthError) {
	return GoogleAuth.Process(args)
}

func (obj *GoogleAuthProvider) Name() string {
	return "Google"
}

func (obj *GoogleAuthProvider) GenerateAuthURL(loginString string) string {
	return GoogleAuth.GenerateAuthURL(loginString)
}
