package user

import (
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/components/user/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
)

func Map(object UserDB) User {
	return User{
		ID:         object.Id,
		Username:   object.Username,
		Password:   object.Password,
		AuthMethod: object.AuthenticationMethod,
		Verified:   object.Verified,
	}
}

func ReverseMap(object User) UserDB {
	return UserDB{
		Id:                   object.ID,
		Username:             object.Username,
		Password:             object.Password,
		AuthenticationMethod: object.AuthMethod,
		Verified:             object.Verified,
	}
}
