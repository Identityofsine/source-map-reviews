package model

type Token struct {
	Id           string `json:"id" db:"id"`
	UserId       int64  `json:"user_id" db:"user_id"`
	AccessToken  string `json:"access_token" db:"access_token"`
	RefreshToken string `json:"refresh_token" db:"refresh_token"`
	ExpiresAt    string `json:"expires_at" db:"expires_at"`
	RefreshedAt  string `json:"refreshed_at" db:"refreshed_at"`
	CreatedAt    string `json:"created_at" db:"created_at"`
}

type SingleToken struct {
	Type       string `json:"type"`
	Token      string `json:"token"`
	Expiration string `json:"expiration"`
}

const (
	TOKEN_TYPE_UNKNOWN = "unknown"
	TOKEN_TYPE_ACCESS  = "access"
	TOKEN_TYPE_REFRESH = "refresh"
)
