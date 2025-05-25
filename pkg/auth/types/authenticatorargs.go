package types

import "github.com/gin-gonic/gin"

type AuthenticatorCallbackFunction *func(c *gin.Context)

type authenticatorArgs struct {
	Keys    map[string]interface{}
	Context *gin.Context
}

type AuthenticatorArgs = *authenticatorArgs

func NewAuthenticatorArgs() AuthenticatorArgs {
	return &authenticatorArgs{
		Keys: make(map[string]interface{}),
	}
}
