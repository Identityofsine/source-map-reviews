package maps

import (
	"github.com/gin-gonic/gin"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/types/router"
)

type route Routeable

func (_ *route) UseRouter(router *gin.RouterGroup) *gin.RouterGroup {
	registerGroup := router.Group("/maps")
	registerGroup.GET("/", GetMapsRoute)
	registerGroup.GET("/:mapName", GetMapRoute)
	registerGroup.POST("/search", SearchMapsRoute)

	// Tag lookup endpoints
	registerGroup.GET("/tags", GetTagsRoute)
	registerGroup.POST("/tags", SaveTagLkRoute)
	registerGroup.DELETE("/tags/:tagLk", DeleteTagLkRoute)

	// Map-tag relationship endpoints
	registerGroup.POST("/:mapName/tags/:tagLk", AddTagToMapRoute)
	registerGroup.DELETE("/:mapName/tags/:tagLk", RemoveTagFromMapRoute)
	registerGroup.PUT("/:mapName/tags", UpdateMapTagsRoute)

	return router
}

var (
	MapRoute = route{}
)
