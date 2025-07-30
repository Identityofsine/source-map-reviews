package repository

import (
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dao"
)

type UserDB struct {
	Id                   int64  `json:"id" db:"id" dao:"omit"`
	Username             string `json:"username" db:"username"`
	Password             string `json:"password" db:"password"`
	AuthenticationMethod string `json:"authentication_method" db:"authentication_method"`
	Verified             bool   `json:"verified" db:"verified" dao:"omit"`
}

// logic
func CreateUser(username, password, authMethod string) db.DatabaseError {

	user := UserDB{
		Username:             username,
		Password:             password,
		AuthenticationMethod: authMethod,
	}

	err := dao.InsertIntoDatabaseByStruct(user)

	return err
}

func CreateUserByUserDb(user *UserDB) db.DatabaseError {

	err := dao.InsertIntoDatabaseByStruct(*user)
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

	rows, err := dao.SelectFromDatabaseByStruct[UserDB](UserDB{}, "username = $1", username)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, db.NewDatabaseError("GetUserByUsername", "User not found", "user-not-found", 404)
	}

	return &(rows)[0], nil
}

func GetUserById(id string) *UserDB {
	rows, err := dao.SelectFromDatabaseByStruct[UserDB](UserDB{}, "id = $1", id)
	if err != nil {
		return nil
	}
	if len(rows) == 0 {
		return nil
	}
	return &(rows)[0]
}
