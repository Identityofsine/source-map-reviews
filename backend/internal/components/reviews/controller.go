package reviews

import (
	"github.com/gin-gonic/gin"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/cookies"
)

// GET /api/v1/reviews/:mapName
func GetMapReviews(c *gin.Context) {
	// Get the map name from the URL parameters
	mapName := c.Param("mapName")
	if mapName == "" {
		c.JSON(exception.CODE_BAD_REQUEST, exception.BadRequest)
	}

	reviews, err := GetReviewsByMapName(mapName)
	if err != nil {
		c.JSON(err.Code, err)
	}

	c.JSON(200, reviews)

	return

}

func SaveMapReview(c *gin.Context) {
	var review MapReview
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(exception.CODE_BAD_REQUEST, exception.BadRequest)
		return
	}

	cookies := cookies.NewCookies(c)
	savedReview, err := SaveReview(review, *cookies)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(200, savedReview)
	return
}
