package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	tokendto "github.com/identityofsine/fofx-go-gin-api-template/api/dto/token"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/components/user/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/types"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/config"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/cookies"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
	"github.com/identityofsine/fofx-go-gin-api-template/util"
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
	generic_map = util.MergeMap(generic_map, secretMap)

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

func CreateAccessToken(userId int64) (*SingleToken, error) {
	config := GetAuthSettings()
	accessTokenExpiration := time.Duration(config.AccessTokenExpiration) * time.Second
	accessToken, err := CreateToken(map[string]any{"user_id": userId, "mode": "access"}, accessTokenExpiration)
	if err != nil {
		return nil, err
	}
	accessToken.Type = TOKEN_TYPE_ACCESS
	return accessToken, nil
}

func CreateRefreshToken(userId int64) (*SingleToken, error) {
	config := GetAuthSettings()
	refreshTokenExpiration := time.Duration(config.RefreshTokenExpiration) * time.Second
	refreshToken, err := CreateToken(map[string]any{"user_id": userId, "mode": "refresh"}, refreshTokenExpiration)
	if err != nil {
		return nil, err
	}
	refreshToken.Type = TOKEN_TYPE_REFRESH
	return refreshToken, nil
}

func CreateLoginToken(userId int64) (*Token, error) {
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

func RenewLoginToken(token Token, user User) (*Token, AuthError) {

	if authError := VerifyUserIsAuthenticated(user, token, TOKEN_TYPE_REFRESH); authError != nil {

		storedlogs.LogError("RenewLoginToken: User is not authenticated or token is invalid", authError)

		return nil, authError
	}

	//delete the old token
	derr := DeleteTokenByRefreshToken(token.RefreshToken)
	if derr != nil {
		err := NewAuthError("RenewLoginToken", "Failed to delete old token", "failed-to-delete-old-token", 500)
		storedlogs.LogError("RenewLoginToken: Failed to delete old token", err)
		return nil, err
	}

	newToken, err := CreateLoginToken(user.ID)
	if err != nil {
		storedlogs.LogError("RenewLoginToken: Failed to create new login token", err)
		return nil, NewAuthError("RenewLoginToken", "Failed to create new login token", "failed-to-create-new-login-token", 500)
	}

	return newToken, nil
}

func StoreTokenIntoCookies(token Token, cookies *cookies.Cookies) error {
	authConfig := GetAuthSettings()

	if err := cookies.SetInt("user_id", token.UserId, authConfig.RefreshTokenExpiration); err != nil {
		return err
	}

	if err := cookies.Set("access_token", token.AccessToken, authConfig.AccessTokenExpiration); err != nil {
		return err
	}

	if err := cookies.Set("refresh_token", token.RefreshToken, authConfig.RefreshTokenExpiration); err != nil {
		return err
	}

	return nil
}

func DeleteTokenInCookies(cookies *cookies.Cookies) error {

	if err := cookies.Delete("user_id"); err != nil {
		return err
	}

	if err := cookies.Delete("access_token"); err != nil {
		return err
	}

	if err := cookies.Delete("refresh_token"); err != nil {
		return err
	}

	return nil
}
