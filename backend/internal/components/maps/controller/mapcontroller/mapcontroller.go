package mapcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps/service/mapgetservice"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
)

func GetMaps(c *gin.Context) {
	storedlogs.LogInfo("GET: /maps")

	// Call the service to get maps
	maps, err := mapgetservice.GetMaps()
	if err != nil {
		storedlogs.LogError("Error getting maps: %v", err)
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, maps)

}
