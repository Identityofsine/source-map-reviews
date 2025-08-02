package user

import (
	"github.com/gin-gonic/gin"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/cookies"
)

func GetUserSelf(c *gin.Context) {

	cookies := cookies.NewCookies(c)
	user, err := GetUserByCookies(cookies)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(200, user)
	return
}
