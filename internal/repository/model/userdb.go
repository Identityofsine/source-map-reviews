package model

import "github.com/identityofsine/fofx-go-gin-api-template/pkg/db"

type UserDB struct {
	Id                   string `json:"id"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	AuthenticationMethod string `json:"authentication_method"`
	Verified             bool   `json:"verified"`
}

// logic
func CreateUser(username, password, authMethod string) db.DatabaseError {
	query := "INSERT INTO users (username, password, authentication_method) VALUES ($1, $2, $3)"

	_, err := db.Query[UserDB](query, username, password, authMethod)

	return err
}

func GetUserByUsername(username string) UserDB {
	query := "SELECT * FROM users WHERE username = $1"
	rows, err := db.Query[UserDB](query, username)
	if err != nil {
		return UserDB{}
	}
	if len(*rows) == 0 {
		return UserDB{}
	}
	return (*rows)[0]
}

func GetUserById(id string) UserDB {
	query := "SELECT * FROM users WHERE id = $1"
	rows, err := db.Query[UserDB](query, id)
	if err != nil {
		return UserDB{}
	}
	if len(*rows) == 0 {
		return UserDB{}
	}
	return (*rows)[0]
}
