package middleware

import "github.com/gin-gonic/gin"

type Middleware interface {
	UseMiddleware(context *gin.Context)
}
