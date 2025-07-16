package cors

import (
	"time"

	corslib "github.com/gin-contrib/cors"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/config"
)

func GetCorsConfigFromYaml(yamlConfig config.CorsConfig) corslib.Config {
	return corslib.Config{
		AllowOrigins:     yamlConfig.AllowOrigins,
		AllowMethods:     yamlConfig.AllowMethods,
		AllowHeaders:     yamlConfig.AllowHeaders,
		ExposeHeaders:    yamlConfig.ExposeHeaders,
		MaxAge:           time.Duration(yamlConfig.MaxAge),
		AllowCredentials: yamlConfig.AllowCredentials,
	}
}
