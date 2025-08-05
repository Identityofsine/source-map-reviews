package maps

import (
	"github.com/gin-gonic/gin"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps/mapsearchform"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/types/routeexception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/cookies"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
)

func GetMapsRoute(c *gin.Context) {
	storedlogs.LogInfo("GET: /maps")

	// Call the service to get maps
	maps, err := GetMaps()
	if err != nil {
		storedlogs.LogError("Error getting maps: %v", err)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(200, maps)

}

func GetMapRoute(c *gin.Context) {
	storedlogs.LogInfo("GET: /maps/:mapName")

	mapName := c.Param("mapName")
	if mapName == "" {
		err := routeexception.NewRouteError(
			nil,
			"Map name is required",
			"map-name-required",
			exception.CODE_BAD_REQUEST,
		)
		storedlogs.LogError("Map name is required", err)
		c.JSON(exception.CODE_BAD_REQUEST, err)
		return
	}

	// Call the service to get a specific map
	mapData, err := GetMap(mapName)
	if err != nil {
		storedlogs.LogError("Error getting map: %v", err)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(200, mapData)
}

func SearchMapsRoute(c *gin.Context) {
	storedlogs.LogInfo("POST: /maps/search")

	var form mapsearchform.MapSearchForm
	if err := c.ShouldBindJSON(&form); err != nil {
		storedlogs.LogError("Error binding search form: %v", err)
		c.JSON(exception.CODE_BAD_REQUEST, gin.H{"error": "Invalid search form"})
		return
	}

	// Call the service to search maps
	maps, err := SearchMaps(form)
	if err != nil {
		storedlogs.LogError("Error searching maps: %v", err)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(200, maps)
}

func GetTagsRoute(c *gin.Context) {
	storedlogs.LogInfo("GET: /maps/tags")

	// Call the service to get map tags
	tags, err := GetTagLks()
	if err != nil {
		storedlogs.LogError("Error getting map tags: %v", err)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(200, tags)
}

// SaveTagLkRoute handles both creating and updating tag lookups
func SaveTagLkRoute(c *gin.Context) {
	storedlogs.LogInfo("POST: /maps/tags")

	var tagLk TagLk
	if err := c.ShouldBindJSON(&tagLk); err != nil {
		storedlogs.LogError("Error binding tag lookup: %v", err)
		c.JSON(exception.CODE_BAD_REQUEST, gin.H{"error": "Invalid tag lookup data"})
		return
	}

	// Get cookies
	cookieJar := cookies.NewCookies(c)

	// Call the service to save tag lookup
	savedTagLk, err := SaveTagLk(tagLk, *cookieJar)
	if err != nil {
		storedlogs.LogError("Error saving tag lookup: %v", err)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(200, savedTagLk)
}

// DeleteTagLkRoute handles deleting tag lookups
func DeleteTagLkRoute(c *gin.Context) {
	storedlogs.LogInfo("DELETE: /maps/tags/:tagLk")

	tagLk := c.Param("tagLk")
	if tagLk == "" {
		err := routeexception.NewRouteError(
			nil,
			"Tag lookup key is required",
			"tag-lk-required",
			exception.CODE_BAD_REQUEST,
		)
		storedlogs.LogError("Tag lookup key is required", err)
		c.JSON(exception.CODE_BAD_REQUEST, err)
		return
	}

	// Get cookies
	cookieJar := cookies.NewCookies(c)

	// Call the service to delete tag lookup
	err := DeleteTagLk(tagLk, *cookieJar)
	if err != nil {
		storedlogs.LogError("Error deleting tag lookup: %v", err)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(200, gin.H{"message": "Tag lookup deleted successfully"})
}

// AddTagToMapRoute handles adding a tag to a map
func AddTagToMapRoute(c *gin.Context) {
	storedlogs.LogInfo("POST: /maps/:mapName/tags/:tagLk")

	mapName := c.Param("mapName")
	tagLk := c.Param("tagLk")

	if mapName == "" {
		err := routeexception.NewRouteError(
			nil,
			"Map name is required",
			"map-name-required",
			exception.CODE_BAD_REQUEST,
		)
		storedlogs.LogError("Map name is required", err)
		c.JSON(exception.CODE_BAD_REQUEST, err)
		return
	}

	if tagLk == "" {
		err := routeexception.NewRouteError(
			nil,
			"Tag lookup key is required",
			"tag-lk-required",
			exception.CODE_BAD_REQUEST,
		)
		storedlogs.LogError("Tag lookup key is required", err)
		c.JSON(exception.CODE_BAD_REQUEST, err)
		return
	}

	// Get cookies
	cookieJar := cookies.NewCookies(c)

	// Call the service to add tag to map
	mapTag, err := AddTagToMap(mapName, tagLk, *cookieJar)
	if err != nil {
		storedlogs.LogError("Error adding tag to map: %v", err)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(200, mapTag)
}

// RemoveTagFromMapRoute handles removing a tag from a map
func RemoveTagFromMapRoute(c *gin.Context) {
	storedlogs.LogInfo("DELETE: /maps/:mapName/tags/:tagLk")

	mapName := c.Param("mapName")
	tagLk := c.Param("tagLk")

	if mapName == "" {
		err := routeexception.NewRouteError(
			nil,
			"Map name is required",
			"map-name-required",
			exception.CODE_BAD_REQUEST,
		)
		storedlogs.LogError("Map name is required", err)
		c.JSON(exception.CODE_BAD_REQUEST, err)
		return
	}

	if tagLk == "" {
		err := routeexception.NewRouteError(
			nil,
			"Tag lookup key is required",
			"tag-lk-required",
			exception.CODE_BAD_REQUEST,
		)
		storedlogs.LogError("Tag lookup key is required", err)
		c.JSON(exception.CODE_BAD_REQUEST, err)
		return
	}

	// Get cookies
	cookieJar := cookies.NewCookies(c)

	// Call the service to remove tag from map
	err := RemoveTagFromMap(mapName, tagLk, *cookieJar)
	if err != nil {
		storedlogs.LogError("Error removing tag from map: %v", err)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(200, gin.H{"message": "Tag removed from map successfully"})
}

// UpdateMapTagsRoute handles updating all tags for a map
func UpdateMapTagsRoute(c *gin.Context) {
	storedlogs.LogInfo("PUT: /maps/:mapName/tags")

	mapName := c.Param("mapName")
	if mapName == "" {
		err := routeexception.NewRouteError(
			nil,
			"Map name is required",
			"map-name-required",
			exception.CODE_BAD_REQUEST,
		)
		storedlogs.LogError("Map name is required", err)
		c.JSON(exception.CODE_BAD_REQUEST, err)
		return
	}

	var request struct {
		Tags []string `json:"tags"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		storedlogs.LogError("Error binding tags update request: %v", err)
		c.JSON(exception.CODE_BAD_REQUEST, gin.H{"error": "Invalid tags data"})
		return
	}

	// Get cookies
	cookieJar := cookies.NewCookies(c)

	// Call the service to update map tags
	updatedTags, err := UpdateMapTags(mapName, request.Tags, *cookieJar)
	if err != nil {
		storedlogs.LogError("Error updating map tags: %v", err)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(200, updatedTags)
}
