package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/config"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
)

func CreateToken(secretMap map[string]any, expiration time.Duration) (*SingleToken, error) {

	config := GetAuthSettings()
	serverName := GetServerDetails().ServerName
	if config.SecretKey == "" {
		storedlogs.LogFatal("CreateToken: SecretKey is empty from GetAuthSettings(). This shouldn't happen and is not allowed. Exiting Application.", nil)
	}

	generic_map := make(map[string]any)
	generic_map["iss"] = serverName
	generic_map["exp"] = time.Now().Add(expiration).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(generic_map))
	token_string, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		return nil, err
	}

	tkn := SingleToken{
		Type:       TOKEN_TYPE_UNKNOWN,
		Token:      token_string,
		Expiration: time.Now().Add(expiration).Format(time.RFC3339),
	}

	return &tkn, nil
}
