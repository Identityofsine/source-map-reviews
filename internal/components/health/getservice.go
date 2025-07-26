package health

import (
	. "github.com/identityofsine/fofx-go-gin-api-template/api/dto/health"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/config"
)

func getHealth() Health {
	// We store this value in the environment variable
	return *MapServerDetailsToHealth(GetServerDetails())
}
