package providers

import (
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth"
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

func (obj *GoogleAuthProvider) authenticate(args AuthenticatorArgs) bool {

	return obj.validate(args)
}
