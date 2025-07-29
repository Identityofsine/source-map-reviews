package bcrypt

import . "golang.org/x/crypto/bcrypt"

func HashString(password string) (string, error) {
	hashedPassword, err := GenerateFromPassword([]byte(password), DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword[:]), nil
}

func CompareHashes(hashedPassword, password string) error {
	return CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
