package providers

import . "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth"

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
	if args["username"] == nil || args["password"] == nil {
		return false
	}

	return true
}

func (obj *InternalAuthProvider) authenticate(args AuthenticatorArgs) bool {
	if valid := obj.validate(args); !valid {
		return false
	}

	return true
}
