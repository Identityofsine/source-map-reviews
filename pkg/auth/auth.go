package auth

import (
	providers "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/providers"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/types"
)

// Authenticator serves as a interface for authentication and for multiple components to handle authentication with the system.
// This isn't something that has to be done for refresh tokens -- that can be handled else where; however anything that is supposed to be an
// "Authenticator" should implement this interface; whether that'd be a login, a refresh token, or something else.
type Authenticator interface {
	Authenticate(args AuthenticatorArgs) bool
	Name() string
}

var (
	authproviders = []Authenticator{
		&providers.GoogleAuthProvider{},
		&providers.InternalAuthProvider{},
	}
)

func GetAuthProviders() []Authenticator {
	return authproviders
}

const (
	AUTH_METHOD_INTERNAL = "internal"
	AUTH_METHOD_GOOGLE   = "google"
)
