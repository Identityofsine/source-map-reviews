package model

import "github.com/identityofsine/fofx-go-gin-api-template/pkg/db"

type UserDB struct {
	Id                   int64  `json:"id"`
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

func CreateUserByUserDb(user *UserDB) db.DatabaseError {
	query := "INSERT INTO users (username, password, authentication_method, verified) VALUES ($1, $2, $3, $4)"

	_, err := db.Query[UserDB](query, user.Username, user.Password, user.AuthenticationMethod, user.Verified)
	if err != nil {
		return err
	}

	nUser, derr := GetUserByUsername(user.Username)
	if derr != nil {
		return derr
	}

	user.Id = nUser.Id

	return err
}

func GetUserByUsername(username string) (*UserDB, db.DatabaseError) {
	query := "SELECT * FROM users WHERE username = $1"
	rows, err := db.Query[UserDB](query, username)
	if err != nil {
		return nil, err
	}
	if len(*rows) == 0 {
		return nil, db.NewDatabaseError("GetUserByUsername", "User not found", "user-not-found", 404)
	}
	return &(*rows)[0], nil
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
