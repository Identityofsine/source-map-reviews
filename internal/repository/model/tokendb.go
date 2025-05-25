package model

import "github.com/identityofsine/fofx-go-gin-api-template/pkg/db"

type TokenDB struct {
	Id           string
	UserId       int64
	AccessToken  string
	RefreshToken string
	ExpiresAt    string
	RefreshedAt  string
	CreatedAt    string
}

func GetTokens() ([]TokenDB, db.DatabaseError) {
	query := "SELECT * FROM public.authtokens"
	rows, err := db.Query[TokenDB](query)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetTokenByUserId(userId string) ([]TokenDB, db.DatabaseError) {
	query := "SELECT * FROM public.authtokens WHERE user_id = $1"
	rows, err := db.Query[TokenDB](query, userId)
	if err != nil {
		return nil, err
	}
	if len(*rows) == 0 {
		return nil, db.NewDatabaseError("GetTokenByUserId", "No tokens found for user", "no-tokens-found", 404)
	}
	return (*rows), nil
}

func GetTokenByAccessToken(accessToken string) (TokenDB, db.DatabaseError) {
	query := "SELECT * FROM public.authtokens WHERE access_token = $1"
	rows, err := db.Query[TokenDB](query, accessToken)
	if err != nil {
		return TokenDB{}, err
	}
	if len(*rows) == 0 {
		return TokenDB{}, nil
	}
	return (*rows)[0], nil
}

func GetTokenByRefreshToken(refreshToken string) (TokenDB, db.DatabaseError) {
	query := "SELECT * FROM public.authtokens WHERE refresh_token = $1"
	rows, err := db.Query[TokenDB](query, refreshToken)
	if err != nil {
		return TokenDB{}, err
	}
	if len(*rows) == 0 {
		return TokenDB{}, nil
	}
	return (*rows)[0], nil
}

func UpdateToken(tokenDB TokenDB) db.DatabaseError {
	query := "UPDATE public.authtokens SET access_token = $1, refresh_token = $2, expires_at = $3, refreshed_at = $4 WHERE user_id = $5"
	_, err := db.Query[TokenDB](query, tokenDB.AccessToken, tokenDB.RefreshToken, tokenDB.ExpiresAt, tokenDB.RefreshedAt, tokenDB.UserId)
	if err != nil {
		return err
	}
	return nil
}

func SaveToken(tokenDB TokenDB) db.DatabaseError {
	query := "INSERT INTO public.authtokens (user_id, access_token, refresh_token, expires_at, refreshed_at, created_at) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := db.Query[TokenDB](query, tokenDB.UserId, tokenDB.AccessToken, tokenDB.RefreshToken, tokenDB.ExpiresAt, tokenDB.RefreshedAt, tokenDB.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTokenById(tokenId string) db.DatabaseError {
	query := "DELETE FROM public.authtokens WHERE id = $1"
	_, err := db.Query[TokenDB](query, tokenId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTokenByRefreshToken(refreshToken string) db.DatabaseError {
	query := "DELETE FROM public.authtokens WHERE refresh_token = $1"
	_, err := db.Query[TokenDB](query, refreshToken)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTokenByUser(userId string) db.DatabaseError {
	query := "DELETE FROM public.authtokens WHERE user_id = $1"
	_, err := db.Query[TokenDB](query, userId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAllTokens() db.DatabaseError {
	query := "DELETE FROM public.authtokens"
	_, err := db.Query[TokenDB](query)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTokensWhen should never be run using any user input; this should be directly controlled by the application logic.
func DeleteTokensWhen(condition string, args ...interface{}) db.DatabaseError {
	query := "DELETE FROM public.authtokens WHERE " + condition
	_, err := db.Query[TokenDB](query, args...)
	if err != nil {
		return err
	}
	return nil
}
