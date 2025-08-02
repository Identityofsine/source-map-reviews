package user

import (
	"strconv"

	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/types/routeexception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/cookies"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper"
)

func GetUserByUserId(userId int64) (*User, routeexception.RouteError) {
	// This function should interact with the user repository to fetch the user by ID.
	// For now, we will return a dummy user for demonstration purposes.
	// In a real application, you would replace this with actual database logic.
	intId := strconv.FormatInt(userId, 10)
	if userDb := GetUserById(intId); userDb.Id == 0 {
		return nil, exception.ResourceNotFound
	} else {
		user := dbmapper.MapDbFields[UserDB, User](*userDb)
		return user, nil
	}
}

// TODO: write RouteError
func GetUserByCookies(cookies *cookies.Cookies) (*User, routeexception.RouteError) {
	// This function should interact with the user repository to fetch the user by cookies.
	// For now, we will return a dummy user for demonstration purposes.
	// In a real application, you would replace this with actual database logic.

	if cookies == nil {
		return nil, routeexception.NewRouteError(nil, "Cookies are nil", "cookies-nil", exception.CODE_BAD_REQUEST)
	}

	// TODO: write constants for cookie keys
	userId, err := cookies.GetInt("user_id")
	if err != nil {
		return nil, routeexception.NewRouteError(err, "Invalid cookies", "invalid-cookies", exception.CODE_BAD_REQUEST)
	}

	return GetUserByUserId(int64(userId))
}
