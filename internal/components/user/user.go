package user

type User struct {
	ID         int64  `json:"id" db:"id"`
	Username   string `json:"username" validate:"required" db:"username"`
	Password   string `json:"password" validate:"required" db:"password" gorm:"-"`
	AuthMethod string `json:"authentication_method" validate:"required" db:"authentication_method"`
	Verified   bool   `json:"verified" db:"verified"`
}
