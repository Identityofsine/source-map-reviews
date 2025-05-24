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
)

func VerifyUserIsAuthenticated(user User, token Token, tokenType string) bool {
	if user.ID == 0 {
		return false
	}

	token_str := token.AccessToken
	if tokenType == TOKEN_TYPE_REFRESH {
		token_str = token.RefreshToken
	}

	claims, err := VerifyToken(token_str)
	if err != nil {
		return false
	}
	if claims["user_id"] == nil {
		return false
	} else if claims["user_id"].(float64) != float64(user.ID) {
		return false
	} else if claims["user_id"].(float64) == float64(user.ID) {
		return true
	}

	return false

}

// VerifyToken solely verifies that the token is valid and not expired.
// This function does not check if the user is authenticated or not; that is, it does not check if the user ID is the same.
func VerifyToken(token_string string) (jwt.MapClaims, AuthError) {

	secret := GetAuthSettings().SecretKey
	//trim spaces
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, NewAuthError("auth", "Error verifying token", err.Error(), 401)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, NewAuthError("auth", "Error verifying token", "Token is invalid", 401)
	}

	/* CHECK FOR EXPIRATION */
	if claims["exp"].(float64) < float64(time.Now().Unix()) {
		return nil, NewAuthError("auth", "Error verifying token", "Token has expired", 401)
	} //gtfo!

	return claims, nil
}

func GetTokenByRefresh(refreshToken string) (*Token, error) {
	tokenDB, derr := GetTokenByRefreshToken(refreshToken)
	if derr != nil {
		return nil, derr
	}
	if tokenDB.Id == "" {
		return nil, nil
	}
	token := tokendto.Map(tokenDB)
	return &token, nil
}
