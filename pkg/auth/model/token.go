package model

type Token struct {
	Id           string `json:"id"`
	UserId       int64  `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    string `json:"expires_at"`
	RefreshedAt  string `json:"refreshed_at"`
	CreatedAt    string `json:"created_at"`
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
