package token

import (
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/model"
)

func Map(db TokenDB) Token {
	return Token{
		Id:           db.Id,
		UserId:       db.UserId,
		AccessToken:  db.AccessToken,
		RefreshToken: db.RefreshToken,
		RefreshedAt:  db.RefreshedAt,
		CreatedAt:    db.CreatedAt,
		ExpiresAt:    db.ExpiresAt,
	}
}

func ReverseMap(token Token) TokenDB {
	return TokenDB{
		Id:           token.Id,
		UserId:       token.UserId,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		RefreshedAt:  token.RefreshedAt,
		CreatedAt:    token.CreatedAt,
		ExpiresAt:    token.ExpiresAt,
	}
}
