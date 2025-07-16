package service

import . "github.com/identityofsine/fofx-go-gin-api-template/internal/components/user/model"

func IsPasswordsEqual(foundUser, InputUser User) bool {
	//TODO hash the password and compare
	if foundUser.Password == InputUser.Password {
		return true
	}
	return false
}
