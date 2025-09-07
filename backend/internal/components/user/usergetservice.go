package user

import (
	"strconv"
	"time"

	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/types/routeexception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/cookies"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper"
)

func GetUserByUserId(userId int64, fullDetails bool) (*User, routeexception.RouteError) {
	// This function should interact with the user repository to fetch the user by ID.
	// For now, we will return a dummy user for demonstration purposes.
	// In a real application, you would replace this with actual database logic.

	var user *User

	intId := strconv.FormatInt(userId, 10)
	if userDb := GetUserById(intId); userDb.Id == 0 {
		return nil, exception.ResourceNotFound
	} else {
		user = dbmapper.MapDbFields[UserDB, User](*userDb)
	}

	UserDetailDb, derr := GetUserDetailByUserId(userId)
	if derr != nil {
		if derr.Code == exception.CODE_RESOURCE_NOT_FOUND {
			// No details found, but user exists - this is acceptable
			return user, nil
		} else {
			// Some other database error occurred
			return nil, routeexception.NewRouteError(derr, "Failed to get user details", "get-user-details-failed", derr.Code)
		}
	}

	UserDetail := dbmapper.MapDbFields[UserDetailDB, UserDetails](*UserDetailDb)

	if !fullDetails {
		UserDetail.Email = ""
		UserDetail.DateOfBirth = time.Now()
		UserDetail.LastName = ""
	}

	user.Details = UserDetail

	return user, nil

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

	return GetUserByUserId(int64(userId), true)
}
