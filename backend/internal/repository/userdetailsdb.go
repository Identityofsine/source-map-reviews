package repository

import (
	"database/sql"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dao"
)

type UserDetailDB struct {
	Id          int64          `json:"id" db:"id" dao:"omit"`
	UserId      int64          `json:"user_id" db:"user_id"`
	Email       sql.NullString `json:"email" db:"email"`
	FirstName   sql.NullString `json:"first_name" db:"first_name"`
	LastName    sql.NullString `json:"last_name" db:"last_name"`
	DateOfBirth sql.NullTime   `json:"date_of_birth" db:"date_of_birth"`
}

func CreateUserDetail(userId int64, email, firstName, lastName, dateOfBirth string) db.DatabaseError {

	UserDetail := UserDetailDB{
		UserId:      userId,
		Email:       sql.NullString{String: email, Valid: email != ""},
		FirstName:   sql.NullString{String: firstName, Valid: firstName != ""},
		LastName:    sql.NullString{String: lastName, Valid: lastName != ""},
		DateOfBirth: sql.NullTime{Valid: false},
	}

	_, err := dao.InsertIntoDatabaseByStruct(UserDetail)

	return err
}

func GetUserDetailByUserId(userId int64) (*UserDetailDB, db.DatabaseError) {

	rows, err := dao.SelectFromDatabaseByStruct[UserDetailDB](UserDetailDB{}, "user_id = $1", userId)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, db.NewDatabaseError("GetUserDetailByUserId", "User details not found", "user-details-not-found", 404)
	}

	return &(rows)[0], nil
}
