package service

import (
	. "github.com/identityofsine/fofx-go-gin-api-template/api/dto/health"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/components/health/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/config"
)

func GetHealth() Health {
	// We store this value in the environment variable
	return *MapServerDetailsToHealth(GetServerDetails())
}
