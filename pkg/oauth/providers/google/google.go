package google

import (
	"context"
	"encoding/json"

	AuthConstants "github.com/identityofsine/fofx-go-gin-api-template/internal/constants/auth"
	userdb "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/model"
	tokenService "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/service"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/types"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/config"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
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
func Process(args AuthenticatorArgs) (*Token, db.DatabaseError) {

	if valid := validate(args); !valid {
		return nil, db.NewDatabaseError("GoogleAuthProvider::Authenticate", "args is nil", "args-nil", 400)
	}

	ifCode := args.Keys["code"]
	code, _ := ifCode.(string)

	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, db.NewDatabaseError("GoogleAuthProvider::Authenticate", "failed to exchange code for token", "failed-to-exchange-code", 400)
	}

	client := oauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, db.NewDatabaseError("GoogleAuthProvider::Authenticate", "failed to get user info from Google", "failed-to-get-user-info", 500)
	}
	defer resp.Body.Close()

	var googleUser struct {
		Email         string `json:"email"`
		Name          string `json:"name"`
		VerifiedEmail bool   `json:"verified_email"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&googleUser); err != nil {
		return nil, db.NewDatabaseError("GoogleAuthProvider::Authenticate", "failed to decode user info", "failed-to-decode-user-info", 500)
	}

	//lets log in the user
	user, derr := userdb.GetUserByUsername(googleUser.Email)
	if derr != nil && derr.Code != 404 {
		return nil, derr
	} else if user == nil {
		user = &userdb.UserDB{
			Username:             googleUser.Email,
			AuthenticationMethod: AuthConstants.AUTHORIZATION_METHOD_GOOGLE,
			Verified:             googleUser.VerifiedEmail,
		}
		derr = userdb.CreateUserByUserDb(user)
		if derr != nil {
			return nil, derr
		}
	}

	// Create a record of their token in the database
	derr = userdb.UpdateOrCreateUserOAuthToken(user.Id, token.AccessToken, token.RefreshToken, AuthConstants.AUTHORIZATION_METHOD_GOOGLE, token.Expiry.Format("2006-01-02 15:04:05"))
	if derr != nil {
		return nil, db.NewDatabaseError("GoogleAuthProvider::Authenticate", "failed to create user OAuth token", "failed-to-create-user-oauth-token", 500)
	}

	// Create a token for the user
	internalToken, err := tokenService.CreateLoginToken(user.Id)
	if err != nil {
		return nil, db.NewDatabaseError("GoogleAuthProvider::Authenticate", "failed to create login token", "failed-to-create-login-token", 500)
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
