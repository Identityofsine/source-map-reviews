package service

import (
	"strconv"

	dto "github.com/identityofsine/fofx-go-gin-api-template/api/dto/user"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/components/user/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/cookies"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

func GetUserByUserId(userId int64) (*User, error) {
	// This function should interact with the user repository to fetch the user by ID.
	// For now, we will return a dummy user for demonstration purposes.
	// In a real application, you would replace this with actual database logic.
	intId := strconv.FormatInt(userId, 10)
	if userDb := GetUserById(intId); userDb.Id == 0 {
		return nil, db.NewDatabaseError("GetUserByUserId", "User not found", "user-not-found", 404)
	} else {
		user := dto.Map(userDb)
		return &user, nil
	}
}

// TODO: write RouteError
func GetUserByCookies(cookies *cookies.Cookies) (*User, error) {
	// This function should interact with the user repository to fetch the user by cookies.
	// For now, we will return a dummy user for demonstration purposes.
	// In a real application, you would replace this with actual database logic.

	if cookies == nil {
		return nil, db.NewDatabaseError("GetUserByCookies", "Cookies are nil", "cookies-nil", 400)
	}

	// TODO: write constants for cookie keys
	userId, err := cookies.GetInt("user_id")
	if err != nil {
		return nil, db.NewDatabaseError("GetUserByCookies", "Invalid cookies", "invalid-cookies", 400)
	}

	return GetUserByUserId(int64(userId))
}
