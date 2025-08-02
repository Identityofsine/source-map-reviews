package model

type Token struct {
	Id           string `json:"id" db:"id"`
	UserId       int64  `json:"userId" db:"user_id"`
	AccessToken  string `json:"accessToken" db:"access_token"`
	RefreshToken string `json:"refreshToken" db:"refresh_token"`
	ExpiresAt    string `json:"expiresAt" db:"expires_at"`
	RefreshedAt  string `json:"refreshedAt" db:"refreshed_at"`
	CreatedAt    string `json:"createdAt" db:"created_at"`
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
