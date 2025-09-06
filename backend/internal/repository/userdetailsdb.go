package repository

import (
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dao"
)

type UserDetailsDB struct {
	Id          int64  `json:"id" db:"id" dao:"omit"`
	UserId      int64  `json:"user_id" db:"user_id"`
	Email       string `json:"email" db:"email"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	DateOfBirth string `json:"date_of_birth" db:"date_of_birth"`
}

func CreateUserDetails(userId int64, email, firstName, lastName, dateOfBirth string) db.DatabaseError {

	userDetails := UserDetailsDB{
		UserId:      userId,
		Email:       email,
		FirstName:   firstName,
		LastName:    lastName,
		DateOfBirth: dateOfBirth,
	}

	_, err := dao.InsertIntoDatabaseByStruct(userDetails)

	return err
}

func GetUserDetailsByUserId(userId int64) (*UserDetailsDB, db.DatabaseError) {

	rows, err := dao.SelectFromDatabaseByStruct[UserDetailsDB](UserDetailsDB{}, "user_id = $1", userId)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, db.NewDatabaseError("GetUserDetailsByUserId", "User details not found", "user-details-not-found", 404)
	}

	return &(rows)[0], nil
}
