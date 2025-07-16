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
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

func VerifyUserIsAuthenticated(user User, token Token, tokenType string) AuthError {
	if user.ID == 0 {
		return NewAuthError("VerifyUserIsAuthenticated", "User is not authenticated", "user-not-authenticated", 401)
	}

	token_str := token.AccessToken
	if tokenType == TOKEN_TYPE_REFRESH {
		token_str = token.RefreshToken
	}

	claims, err := VerifyToken(token_str)
	if err != nil {
		return err
	}

	var dbToken TokenDB
	var derr db.DatabaseError

	if tokenType == TOKEN_TYPE_ACCESS {
		dbToken, derr = GetTokenByAccessToken(token_str)
	} else if tokenType == TOKEN_TYPE_REFRESH {
		dbToken, derr = GetTokenByRefreshToken(token_str)
	} else {
		return NewAuthError("VerifyUserIsAuthenticated", "Invalid token type", "invalid-token-type", 400)
	}

	if derr != nil {
		return NewAuthError("VerifyUserIsAuthenticated", derr.Message, derr.Err, derr.Code)
	} else if dbToken.Id == "" {
		return NewAuthError("VerifyUserIsAuthenticated", "Token not found in database", "token-not-found", 404)
	}

	if claims["user_id"] == nil {
		return NewAuthError("VerifyUserIsAuthenticated", "Token does not contain user_id", "token-missing-user-id", 401)
	} else if claims["user_id"].(float64) != float64(user.ID) {
		return NewAuthError("VerifyUserIsAuthenticated", "User ID in token does not match authenticated user", "user-id-mismatch", 401)
	} else if claims["user_id"].(float64) == float64(user.ID) {
		return nil
	}

	return NewAuthError("VerifyUserIsAuthenticated", "User ID in token does not match authenticated user", "user-id-mismatch", 401)
}

// VerifyToken solely verifies that the token is valid and not expired.
// This function does not check if the user is authenticated or not; that is, it does not check if the user ID is the same.
func VerifyToken(token_string string) (jwt.MapClaims, AuthError) {

	secret := []byte(GetAuthSettings().SecretKey)
	//trim spaces
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, NewAuthError("auth", err.Error(), "parsing-failed", 401)
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

func GetTokenFromCookies(cookies *cookies.Cookies) (*Token, AuthError) {
	if cookies == nil {
		return nil, NewAuthError("GetTokensFromCookies", "Cookies are nil", "cookies-nil", 400)
	}

	accessToken, err := cookies.Get("access_token")
	if err != nil || accessToken == "" {
		return nil, NewAuthError("GetTokensFromCookies", "Access token is required", "access-token-required", 401)
	}

	refreshToken, err := cookies.Get("refresh_token")
	if err != nil || refreshToken == "" {
		return nil, NewAuthError("GetTokensFromCookies", "Refresh token is required", "refresh-token-required", 401)
	}

	userId, err := cookies.GetInt("user_id")
	if err != nil || userId <= 0 {
		return nil, NewAuthError("GetTokensFromCookies", "User ID is required in cookies", "user-id-required", 401)
	}

	token := &Token{
		UserId:       int64(userId),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return token, nil
}

func GetAccessTokenFromHeader(header string) (string, AuthError) {
	if header == "" {
		return "", NewAuthError("GetAccessTokenFromHeader", "Authorization header is required", "authorization-header-required", 400)
	}
	if len(header) < 7 || header[:7] != "Bearer " {
		return "", NewAuthError("GetAccessTokenFromHeader", "Authorization header must start with 'Bearer '", "invalid-authorization-header", 400)
	}
	return header[7:], nil
}
