package mapcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps/model/mapsearchform"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps/service/mapgetservice"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/types/routeexception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
)

func GetMaps(c *gin.Context) {
	storedlogs.LogInfo("GET: /maps")

	// Call the service to get maps
	maps, err := mapgetservice.GetMaps()
	if err != nil {
		storedlogs.LogError("Error getting maps: %v", err)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(200, maps)

}

func GetMap(c *gin.Context) {
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
	mapData, err := mapgetservice.GetMap(mapName)
	if err != nil {
		storedlogs.LogError("Error getting map: %v", err)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(200, mapData)
}

func SearchMaps(c *gin.Context) {
	storedlogs.LogInfo("POST: /maps/search")

	var form mapsearchform.MapSearchForm
	if err := c.ShouldBindJSON(&form); err != nil {
		storedlogs.LogError("Error binding search form: %v", err)
		c.JSON(exception.CODE_BAD_REQUEST, gin.H{"error": "Invalid search form"})
		return
	}

	// Call the service to search maps
	maps, err := mapgetservice.SearchMaps(form)
	if err != nil {
		storedlogs.LogError("Error searching maps: %v", err)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(200, maps)
}
