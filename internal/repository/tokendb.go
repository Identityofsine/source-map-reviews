package repository

import (
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dao"
)

type AuthTokenDB struct {
	Id           string `db:"id" dao:"omit"`
	UserId       int64  `db:"user_id"`
	AccessToken  string `db:"access_token"`
	RefreshToken string `db:"refresh_token"`
	ExpiresAt    string `db:"expires_at"`
	RefreshedAt  string `db:"refreshed_at"`
	CreatedAt    string `db:"created_at"`
}

// GetTokens retrieves all tokens from the database.
func GetTokens() ([]AuthTokenDB, db.DatabaseError) {
	return dao.SelectFromDatabaseByStruct(
		AuthTokenDB{},
		"")
}

func GetTokenByUserId(userId string) ([]AuthTokenDB, db.DatabaseError) {
	rows, err := dao.SelectFromDatabaseByStruct(AuthTokenDB{}, "user_id = $1", userId)

	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, db.NewDatabaseError("GetTokenByUserId", "No tokens found for user", "no-tokens-found", 404)
	}

	return rows, nil

}

func GetTokenByAccessToken(accessToken string) (*AuthTokenDB, db.DatabaseError) {

	rows, err := dao.SelectFromDatabaseByStruct(AuthTokenDB{}, "access_token = $1", accessToken)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, exception.ResourceNotFoundDatabase
	}

	return &rows[0], nil
}

func GetTokenByRefreshToken(refreshToken string) (*AuthTokenDB, db.DatabaseError) {

	rows, err := dao.SelectFromDatabaseByStruct(AuthTokenDB{}, "refresh_token = $1", refreshToken)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, exception.ResourceNotFoundDatabase
	}

	return &rows[0], nil
}

func UpdateToken(tokenDB AuthTokenDB) db.DatabaseError {
	query := "UPDATE public.authtokens SET access_token = $1, refresh_token = $2, expires_at = $3, refreshed_at = $4 WHERE user_id = $5"
	_, err := db.Query[AuthTokenDB](query, tokenDB.AccessToken, tokenDB.RefreshToken, tokenDB.ExpiresAt, tokenDB.RefreshedAt, tokenDB.UserId)
	if err != nil {
		return err
	}
	return nil
}

func SaveToken(tokenDB AuthTokenDB) db.DatabaseError {
	err := dao.InsertIntoDatabaseByStruct(tokenDB)

	return err
}

func DeleteTokenById(tokenId string) db.DatabaseError {
	query := "DELETE FROM public.authtokens WHERE id = $1"
	_, err := db.Query[AuthTokenDB](query, tokenId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTokenByRefreshToken(refreshToken string) db.DatabaseError {
	query := "DELETE FROM public.authtokens WHERE refresh_token = $1"
	_, err := db.Query[AuthTokenDB](query, refreshToken)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTokenByUser(userId string) db.DatabaseError {
	query := "DELETE FROM public.authtokens WHERE user_id = $1"
	_, err := db.Query[AuthTokenDB](query, userId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAllTokens() db.DatabaseError {
	query := "DELETE FROM public.authtokens"
	_, err := db.Query[AuthTokenDB](query)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTokensWhen should never be run using any user input; this should be directly controlled by the application logic.
func DeleteTokensWhen(condition string, args ...interface{}) db.DatabaseError {
	query := "DELETE FROM public.authtokens WHERE " + condition
	_, err := db.Query[AuthTokenDB](query, args...)
	if err != nil {
		return err
	}
	return nil
}
