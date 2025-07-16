package providers

import (
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/types"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	GoogleAuth "github.com/identityofsine/fofx-go-gin-api-template/pkg/oauth/providers/google"
)

type GoogleAuthProvider struct {
}

// TODO: replace db.DatabaseError with a more specific error type, with actionable error messages
func (obj *GoogleAuthProvider) Authenticate(args AuthenticatorArgs) (*Token, db.DatabaseError) {
	return GoogleAuth.Process(args)
}

func (obj *GoogleAuthProvider) Name() string {
	return "Google"
}

func (obj *GoogleAuthProvider) GenerateAuthURL(loginString string) string {
	return GoogleAuth.GenerateAuthURL(loginString)
}
