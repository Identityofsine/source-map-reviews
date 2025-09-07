package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/types/routeexception"
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

func GetUserNameById(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(400, "userId is required")
		return
	}

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(400, exception.BadRequest)
		return
	}

	user, derr := GetUserByUserId(int64(userIdInt), false)
	if err != nil {
		c.JSON(derr.Code, err)
		return
	}
	if user == nil {
		c.JSON(404, exception.ResourceNotFound)
		return
	}

	if user.Details != nil {
		if user.Details.FirstName == "" {
			c.JSON(404, routeexception.NewRouteError(
				nil,
				"User has no details set",
				"user-has-no-details",
				exception.CODE_RESOURCE_NOT_FOUND,
			))
			return
		}
	} else {
		c.JSON(404, routeexception.NewRouteError(
			nil,
			"User has no details set",
			"user-has-no-details",
			exception.CODE_RESOURCE_NOT_FOUND,
		))
		return
	}

	c.JSON(200, user)
	return
}
