package internal

import "github.com/gin-gonic/gin"

type Router interface {
	UseRouter(parent *gin.RouterGroup) *gin.RouterGroup
}

type Routeable struct{}
