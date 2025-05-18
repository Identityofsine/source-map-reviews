package auth

type AuthenticatorArgs map[string]interface{}

// Authenticator serves as a interface for authentication and for multiple components to handle authentication with the system.
// This isn't something that has to be done for refresh tokens -- that can be handled else where; however anything that is supposed to be an
// "Authenticator" should implement this interface; whether that'd be a login, a refresh token, or something else.
type Authenticator interface {
	authenticate(args AuthenticatorArgs) bool
}
