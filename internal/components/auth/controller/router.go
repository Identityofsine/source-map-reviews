package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/types/router"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/auth"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/config"
)

type route Routeable

func (_ *route) UseRouter(router *gin.RouterGroup) *gin.RouterGroup {
	authProviders := auth.GetAuthProviders()
	authRoute := router.Group("/auth")
	loginGroup := authRoute.Group("/login")
	{
		for _, provider := range authProviders {
			provider := provider
			providerName := strings.ToLower(provider.Name())
			oAuth, ok := provider.(auth.OAuthAuthenticator)
			if ok {
				loginGroup.GET("/"+providerName, func(c *gin.Context) {
					GenericAuthHandler(provider, c)
				})
				loginGroup.GET("/"+providerName+"/redirect", func(c *gin.Context) {
					redirectURL := oAuth.GenerateAuthURL(serverConfig.WebServerConfig.URI + "api/v1/auth/" + providerName)
					c.Redirect(302, redirectURL)
				})
			} else {
				loginGroup.POST("/"+providerName, func(c *gin.Context) {
					GenericAuthHandler(provider, c)
				})
			}
		}
	}

	//refresh
	authRoute.GET("/refresh", RefershTokenHandler)

	//logout
	authRoute.GET("/logout", LogoutHandler)

	return router
}

var (
	AuthRoute    = route{}
	serverConfig = config.GetServerDetails()
)
