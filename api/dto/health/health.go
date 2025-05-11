package health

import (
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/components/health/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/config"
)

func MapServerDetailsToHealth(serverDetails *ServerDetails) *Health {
	return &Health{
		ServerName:  serverDetails.ServerName,
		BuildDate:   serverDetails.BuildDate,
		Version:     serverDetails.Version,
		Commit:      serverDetails.Commit,
		Branch:      serverDetails.Branch,
		Environment: serverDetails.Environment,
	}
}
