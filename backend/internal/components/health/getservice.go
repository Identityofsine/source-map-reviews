package health

import (
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/config"
)

func GetHealth() Health {
	return *MapServerDetailsToHealth(*GetServerDetails())
}

func MapServerDetailsToHealth(serverDetails ServerDetails) *Health {
	return &Health{
		Version:     serverDetails.Version,
		Commit:      serverDetails.Commit,
		BuildDate:   serverDetails.BuildDate,
		Environment: serverDetails.Environment,
	}
}
