package route

import (
	"github.com/gin-gonic/gin"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/components/user"

	. "github.com/identityofsine/fofx-go-gin-api-template/internal/types/router"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/middlewares"
)

type route Routeable

//this directory contains the routes responsible for handling the requests
//of the health component of this web application

func (_ *route) UseRouter(router *gin.RouterGroup) *gin.RouterGroup {
	g := router.Group("/user")
	g.GET("/:userId", GetUserNameById)

	g.Use(middlewares.UseAuthenticationEnforcementMiddleware().Middleware)

	g.GET("/me", GetUserSelf)
	return router
}

var (
	UserRoute = route{}
)
