package model

type User struct {
	ID         int64  `json:"id"`
	Username   string `json:"username" validate:"required"`
	Password   string `json:"password" validate:"required"`
	AuthMethod string `json:"authentication_method" validate:"required"`
	Verified   bool   `json:"verified"`
}
