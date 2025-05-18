package model

import "github.com/identityofsine/fofx-go-gin-api-template/pkg/db"

type TokenDB struct {
	Id           string
	UserId       string
	AccessToken  string
	RefreshToken string
	ExpiresAt    string
	RefreshedAt  string
	CreatedAt    string
}

func GetTokens() ([]TokenDB, error) {
	query := "SELECT * FROM public.tokens"
	rows, err := db.Query[TokenDB](query)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetTokenByUserId(userId string) (TokenDB, error) {
	query := "SELECT * FROM public.tokens WHERE user_id = $1"
	rows, err := db.Query[TokenDB](query, userId)
	if err != nil {
		return TokenDB{}, err
	}
	if len(*rows) == 0 {
		return TokenDB{}, nil
	}
	return (*rows)[0], nil
}

func UpdateToken(tokenDB TokenDB) error {
	query := "UPDATE public.tokens SET access_token = $1, refresh_token = $2, expires_at = $3, refreshed_at = $4 WHERE user_id = $5"
	_, err := db.Query[TokenDB](query, tokenDB.AccessToken, tokenDB.RefreshToken, tokenDB.ExpiresAt, tokenDB.RefreshedAt, tokenDB.UserId)
	if err != nil {
		return err
	}
	return nil
}

func SaveToken(tokenDB TokenDB) error {
	query := "INSERT INTO public.tokens (user_id, access_token, refresh_token, expires_at, refreshed_at, created_at) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := db.Query[TokenDB](query, tokenDB.UserId, tokenDB.AccessToken, tokenDB.RefreshToken, tokenDB.ExpiresAt, tokenDB.RefreshedAt, tokenDB.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func DeleteToken(userId string) error {
	query := "DELETE FROM public.tokens WHERE user_id = $1"
	_, err := db.Query[TokenDB](query, userId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAllTokens() error {
	query := "DELETE FROM public.tokens"
	_, err := db.Query[TokenDB](query)
	if err != nil {
		return err
	}
	return nil
}
