package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	tokendto "github.com/identityofsine/fofx-go-gin-api-template/api/dto/token"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
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

func CreateAccessToken(userId string) (*SingleToken, error) {
	config := GetAuthSettings()
	accessTokenExpiration := time.Duration(config.AccessTokenExpiration) * time.Second
	accessToken, err := CreateToken(map[string]any{"user_id": userId}, accessTokenExpiration)
	if err != nil {
		return nil, err
	}
	accessToken.Type = TOKEN_TYPE_ACCESS
	return accessToken, nil
}

func CreateRefreshToken(userId string) (*SingleToken, error) {
	config := GetAuthSettings()
	refreshTokenExpiration := time.Duration(config.RefreshTokenExpiration) * time.Second
	refreshToken, err := CreateToken(map[string]any{"user_id": userId}, refreshTokenExpiration)
	if err != nil {
		return nil, err
	}
	refreshToken.Type = TOKEN_TYPE_REFRESH
	return refreshToken, nil
}

func CreateLoginToken(userId string) (*Token, error) {
	accessToken, err := CreateAccessToken(userId)
	if err != nil {
		return nil, err
	}
	refreshToken, err := CreateRefreshToken(userId)
	if err != nil {
		return nil, err
	}

	loginToken := Token{
		Id:           "",
		UserId:       userId,
		AccessToken:  accessToken.Token,
		RefreshToken: refreshToken.Token,
		ExpiresAt:    refreshToken.Expiration,
		RefreshedAt:  time.Now().Format(time.RFC3339),
		CreatedAt:    time.Now().Format(time.RFC3339),
	}

	loginTokenDB := tokendto.ReverseMap(loginToken)
	derr := SaveToken(loginTokenDB)
	if derr != nil {
		return nil, derr
	}

	return &loginToken, nil
}
