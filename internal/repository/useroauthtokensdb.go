package repository

import "github.com/identityofsine/fofx-go-gin-api-template/pkg/db"

type UserOAuthTokenDB struct {
	UserId       int64  `json:"user_id" db:"user_id"`             // User ID associated with the OAuth token
	AccessToken  string `json:"access_token" db:"access_token"`   // OAuth access token
	RefreshToken string `json:"refresh_token" db:"refresh_token"` // OAuth refresh token, if applicable
	Source       string `json:"source" db:"source"`               // e.g., "google", "github", etc.
	CreatedAt    string `json:"created_at" db:"created_at"`       // Timestamp when the token was created
	ExpiresAt    string `json:"expires_at" db:"expires_at"`       // Timestamp when the token expires
}

func CreateUserOAuthToken(userId int64, accessToken, refreshToken, source, expires_at string) db.DatabaseError {

	query := "INSERT INTO user_oauth_tokens (user_id, access_token, refresh_token, source, expires_at) VALUES ($1, $2, $3, $4, $5)"

	_, err := db.Query[UserOAuthTokenDB](query, userId, accessToken, refreshToken, source, expires_at)

	return err
}

func UpdateOrCreateUserOAuthToken(userId int64, accessToken, refreshToken, source, expires_at string) db.DatabaseError {
	query := `
		INSERT INTO user_oauth_tokens (user_id, access_token, refresh_token, source, expires_at)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (user_id, source) DO UPDATE
		SET access_token = EXCLUDED.access_token,
			refresh_token = EXCLUDED.refresh_token,
			expires_at = EXCLUDED.expires_at
	`
	_, err := db.Query[UserOAuthTokenDB](query, userId, accessToken, refreshToken, source, expires_at)
	return err
}

func GetUserOAuthTokenByUserIdAndSource(userId int64, source string) (*UserOAuthTokenDB, db.DatabaseError) {
	query := "SELECT * FROM user_oauth_tokens WHERE user_id = $1 AND source = $2"
	rows, err := db.Query[UserOAuthTokenDB](query, userId, source)
	if err != nil {
		return nil, err
	}
	if len(*rows) == 0 {
		return nil, db.NewDatabaseError("GetUserOAuthTokenByUserIdAndSource", "OAuth token not found", "oauth-token-not-found", 404)
	}
	return &(*rows)[0], nil
}
