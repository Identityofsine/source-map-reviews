package user

func IsPasswordsEqual(foundUser, InputUser User) bool {
	//TODO hash the password and compare
	if foundUser.Password == InputUser.Password {
		return true
	}
	return false
}
