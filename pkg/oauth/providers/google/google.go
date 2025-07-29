package google

import (
	"context"
	"encoding/json"

	AuthConstants "github.com/identityofsine/fofx-go-gin-api-template/internal/constants/auth"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	userdb "github.com/identityofsine/fofx-go-gin-api-template/internal/repository"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/authtypes"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/authtypes"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/model"
	tokenService "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/service"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	serverConfig     = config.GetServerDetails()
	authConfig       = config.GetAuthSettings()
	googleAuthConfig = authConfig.GoogleAuthSecrets
	oauthConfig      = &oauth2.Config{
		RedirectURL:  serverConfig.WebServerConfig.URI + "api/v1/auth/login/google",
		ClientID:     googleAuthConfig.ClientID,
		ClientSecret: googleAuthConfig.ClientSecret,
		Scopes:       googleAuthConfig.Scopes,
		Endpoint:     google.Endpoint,
	}
)

// move this to validators?
func validate(args AuthenticatorArgs) bool {
	//check if args is nil or what we expect from a google auth request
	if args == nil {
		return false
	}

	if args.Keys == nil || args.Keys["code"] == nil {
		return false
	}

	// Check if the code is a string
	if _, ok := args.Keys["code"].(string); !ok {
		return false
	}

	return true
}

// TODO: replace db.DatabaseError with a more specific error type, with actionable error messages
func Process(args AuthenticatorArgs) (*Token, authtypes.AuthError) {

	if valid := validate(args); !valid {
		return nil, authtypes.NewAuthError(
			"GoogleAuthProvider::Authenticate",
			"invalid-credentials",
			"invalid-credentials",
			exception.CODE_BAD_REQUEST,
		)
	}

	ifCode := args.Keys["code"]
	code, _ := ifCode.(string)

	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, authtypes.NewAuthError(
			"GoogleAuthProvider::Authenticate",
			"failed to exchange code for token",
			"failed-to-exchange-code",
			exception.CODE_UNAUTHORIZED,
		)
	}

	client := oauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, authtypes.NewAuthError(
			"GoogleAuthProvider::Authenticate",
			"failed to get user info from Google",
			"failed-to-get-user-info",
			exception.CODE_UNAUTHORIZED,
		)
	}
	defer resp.Body.Close()

	var googleUser struct {
		Email         string `json:"email"`
		Name          string `json:"name"`
		VerifiedEmail bool   `json:"verified_email"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&googleUser); err != nil {
		return nil, authtypes.NewAuthError(
			"GoogleAuthProvider::Authenticate",
			"failed to decode user info from Google",
			"failed-to-decode-user-info",
			exception.CODE_UNAUTHORIZED,
		)
	}

	//lets log in the user
	user, derr := userdb.GetUserByUsername(googleUser.Email)
	if derr != nil && derr.Code != 404 {
		return nil, authtypes.NewAuthError(
			"GoogleAuthProvider::Authenticate",
			derr.Message,
			"failed-to-get-user",
			derr.Code,
		)
	} else if user == nil {
		user = &userdb.UserDB{
			Username:             googleUser.Email,
			AuthenticationMethod: AuthConstants.AUTHORIZATION_METHOD_GOOGLE,
			Verified:             googleUser.VerifiedEmail,
		}
		derr = userdb.CreateUserByUserDb(user)
		if derr != nil {
			return nil, authtypes.NewAuthError(
				"GoogleAuthProvider::Authenticate",
				derr.Message,
				"failed-to-create-user",
				derr.Code,
			)
		}
	}

	// Create a record of their token in the database
	derr = userdb.UpdateOrCreateUserOAuthToken(user.Id, token.AccessToken, token.RefreshToken, AuthConstants.AUTHORIZATION_METHOD_GOOGLE, token.Expiry.Format("2006-01-02 15:04:05"))
	if derr != nil {
		return nil, authtypes.NewAuthError(
			"GoogleAuthProvider::Authenticate",
			derr.Message,
			"failed-to-create-oauth-token",
			derr.Code,
		)
	}

	// Create a token for the user
	internalToken, err := tokenService.CreateLoginToken(user.Id)
	if err != nil {
		return nil, authtypes.NewAuthError(
			"GoogleAuthProvider::Authenticate",
			"failed to create login token",
			"failed-to-create-login-token",
			exception.CODE_INTERNAL_SERVER_ERROR,
		)
	}

	return internalToken, nil
}

func GenerateAuthURL(originalPath string) string {
	if oauthConfig == nil {
		return ""
	}
	oauth2.SetAuthURLParam("prompt", "consent")
	return oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
}
