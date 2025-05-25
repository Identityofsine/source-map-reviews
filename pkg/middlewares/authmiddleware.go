package middlewares

import (
	"github.com/gin-gonic/gin"
	UserService "github.com/identityofsine/fofx-go-gin-api-template/internal/components/user/service"
	Token "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/model"
	AuthService "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/service"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/types"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/cookies"
)

type authmiddleware struct {
	Middleware gin.HandlerFunc
}

// CreateCors is a constructor function that creates a new CORS middleware; taking in a list of valid origins
func UseAuthenticationEnforcementMiddleware() *authmiddleware {
	return &authmiddleware{
		Middleware: validateToken,
	}
}

func validateToken(c *gin.Context) {
	cookies := cookies.NewCookies(c)
	// Check if the user is authenticated
	var aerr AuthError
	accessToken := c.GetHeader("Authorization")
	accessToken, aerr = AuthService.GetAccessTokenFromHeader(accessToken)
	if accessToken == "" || aerr != nil {
		c.AbortWithStatusJSON(aerr.Code, aerr)
		return
	}

	//get the user
	user, err := UserService.GetUserByCookies(cookies)
	if err != nil || user == nil {
		aerr = NewAuthError("validateToken", "Not authenticated", "user-not-authed", 401)
		c.AbortWithStatusJSON(aerr.Code, aerr)
		return
	}

	token, aerr := AuthService.GetTokenFromCookies(cookies)
	if aerr != nil {
		c.AbortWithStatusJSON(aerr.Code, aerr)
		return
	}

	token.AccessToken = accessToken

	// Verify the token
	aerr = AuthService.VerifyUserIsAuthenticated(*user, *token, Token.TOKEN_TYPE_ACCESS)
	if aerr != nil {
		c.AbortWithStatusJSON(aerr.Code, aerr)
		return
	}

	c.Next()
}
