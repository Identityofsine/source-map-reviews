package auth

import (
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/model"
	providers "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/providers"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/types"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

// Authenticator serves as a interface for authentication and for multiple components to handle authentication with the system.
// This isn't something that has to be done for refresh tokens -- that can be handled else where; however anything that is supposed to be an
// "Authenticator" should implement this interface; whether that'd be a login, a refresh token, or something else.
type Authenticator interface {
	Authenticate(args AuthenticatorArgs) (*Token, db.DatabaseError) // --> /login/Name()
	Name() string
}

type OAuthAuthenticator interface {
	Authenticator
	GenerateAuthURL(loginPath string) string // --> /auth/Name()/generate : redirects to the auth provider's login page : redirects to /login/Name()
}

var (
	authproviders = []Authenticator{
		&providers.GoogleAuthProvider{},
		&providers.InternalAuthProvider{},
	}
	oauthproviders = []OAuthAuthenticator{
		&providers.GoogleAuthProvider{},
	}
)

func GetAuthProviders() []Authenticator {
	return authproviders
}

func GetOAuthProviders() []OAuthAuthenticator {
	return oauthproviders
}
