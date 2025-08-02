package reviews

import (
	"github.com/gin-gonic/gin"

	. "github.com/identityofsine/fofx-go-gin-api-template/internal/types/router"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/middlewares"
)

type route Routeable

func (_ *route) UseRouter(router *gin.RouterGroup) *gin.RouterGroup {

	g := router.Group("/reviews")
	g.GET("/:mapName", GetMapReviews)

	// Use the authentication enforcement middleware for posting
	g.Use(middlewares.UseAuthenticationEnforcementMiddleware().Middleware)
	g.POST("/:mapName", SaveMapReview)

	return router
}

var (
	ReviewsRoute = route{}
)
