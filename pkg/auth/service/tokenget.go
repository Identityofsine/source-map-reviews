package service

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/components/user"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/authtypes"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/config"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/cookies"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper"
)

func VerifyUserIsAuthenticated(user User, token Token, tokenType string) AuthError {
	if user.ID == 0 {
		return NewAuthError("VerifyUserIsAuthenticated", "User is not authenticated", "user-not-authenticated", exception.CODE_UNAUTHORIZED)
	}

	token_str := token.AccessToken
	if tokenType == TOKEN_TYPE_REFRESH {
		token_str = token.RefreshToken
	}

	claims, err := VerifyToken(token_str)
	if err != nil {
		return err
	}

	var dbToken *TokenDB
	var derr db.DatabaseError

	if tokenType == TOKEN_TYPE_ACCESS {
		dbToken, derr = GetTokenByAccessToken(token_str)
	} else if tokenType == TOKEN_TYPE_REFRESH {
		dbToken, derr = GetTokenByRefreshToken(token_str)
	} else {
		return NewAuthError("VerifyUserIsAuthenticated", "Invalid token type", "invalid-token-type", exception.CODE_BAD_REQUEST)
	}

	if derr != nil {
		return NewAuthError("VerifyUserIsAuthenticated", derr.Message, derr.Err, derr.Code)
	} else if dbToken.Id == "" {
		return NewAuthError("VerifyUserIsAuthenticated", "Token not found in database", "token-not-found", 404)
	}

	if claims["user_id"] == nil {
		return NewAuthError("VerifyUserIsAuthenticated", "Token does not contain user_id", "token-missing-user-id", exception.CODE_UNAUTHORIZED)
	} else if claims["user_id"].(float64) != float64(user.ID) {
		return NewAuthError("VerifyUserIsAuthenticated", "User ID in token does not match authenticated user", "user-id-mismatch", exception.CODE_UNAUTHORIZED)
	} else if claims["user_id"].(float64) == float64(user.ID) {
		return nil
	}

	return NewAuthError("VerifyUserIsAuthenticated", "User ID in token does not match authenticated user", "user-id-mismatch", exception.CODE_UNAUTHORIZED)
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
		if strings.Contains(err.Error(), "token is expired") {
			return nil, exception.TokenExpired
		}
		return nil, NewAuthError("auth", err.Error(), "parsing-failed", exception.CODE_UNAUTHORIZED)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, NewAuthError("auth", "Error verifying token", "Token is invalid", exception.CODE_UNAUTHORIZED)
	}

	/* CHECK FOR EXPIRATION */
	if claims["exp"].(float64) < float64(time.Now().Unix()) {
		return nil, exception.TokenExpired
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
	token := dbmapper.MapDbFields[TokenDB, Token](*tokenDB)
	return token, nil
}

func GetTokenFromCookies(cookies *cookies.Cookies) (*Token, AuthError) {
	if cookies == nil {
		return nil, NewAuthError("GetTokensFromCookies", "Cookies are nil", "cookies-nil", exception.CODE_BAD_REQUEST)
	}

	accessToken, err := cookies.Get("access_token")
	if err != nil || accessToken == "" {
		return nil, NewAuthError("GetTokensFromCookies", "Access token is required", "access-token-required", exception.CODE_UNAUTHORIZED)
	}

	refreshToken, err := cookies.Get("refresh_token")
	if err != nil || refreshToken == "" {
		return nil, NewAuthError("GetTokensFromCookies", "Refresh token is required", "refresh-token-required", exception.CODE_UNAUTHORIZED)
	}

	userId, err := cookies.GetInt("user_id")
	if err != nil || userId <= 0 {
		return nil, exception.TokenExpired
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
		return "", NewAuthError("GetAccessTokenFromHeader", "Authorization header is required", "authorization-header-required", exception.CODE_BAD_REQUEST)
	}
	if len(header) < 7 || header[:7] != "Bearer " {
		return "", NewAuthError("GetAccessTokenFromHeader", "Authorization header must start with 'Bearer '", "invalid-authorization-header", exception.CODE_BAD_REQUEST)
	}
	return header[7:], nil
}
