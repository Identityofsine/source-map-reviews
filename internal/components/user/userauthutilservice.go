package user

import "github.com/identityofsine/fofx-go-gin-api-template/pkg/bcrypt"

func IsPasswordsEqual(foundUser, InputUser User) bool {
	//TODO hash the password and compare
	if err := bcrypt.CompareHashes(foundUser.Password, InputUser.Password); err == nil {
		return true
	}
	return false
}
