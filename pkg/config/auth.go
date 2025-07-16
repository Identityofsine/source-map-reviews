package config

import (
	"fmt"
	"os"
)

//This file contains the logic that is responsible for details
//of the server, such as the version, branch, commit, and build date.
//This is primarily used in the health component of the web application

type GoogleAuthSecrets struct {
	ClientID     string `yaml:"clientId"`
	ClientSecret string
	Scopes       []string `yaml:"scopes"`
}

type AuthSettings struct {
	ConfigFile
	AccessTokenExpiration  int               `yaml:"accessTokenExpiration"`
	RefreshTokenExpiration int               `yaml:"refreshTokenExpiration"`
	GoogleAuthSecrets      GoogleAuthSecrets `yaml:"google"`
	SecretKey              string
}

var (
	cachedConfig *AuthSettings
)

func GetAuthSettings() *AuthSettings {
	if cachedConfig != nil {
		return cachedConfig
	}
	if config, err := loadConfig[*AuthSettings]("auth"); err == nil {
		config.GoogleAuthSecrets.ClientSecret = os.Getenv("GOOGLE_SECRET_KEY")
		config.SecretKey = getSecretKey()
		cachedConfig = config
		return config
	} else {
		fmt.Printf("Error loading auth config: %v\n", err)
		return &AuthSettings{
			AccessTokenExpiration:  3600,
			RefreshTokenExpiration: 604800,
			SecretKey:              getSecretKey(),
		}
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	return secretKey
}
