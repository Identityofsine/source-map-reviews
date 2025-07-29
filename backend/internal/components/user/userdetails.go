package user

import "time"

type UserDetails struct {
	ID          int64     `json:"id" db:"id"`
	UserId      int64     `json:"user_id" db:"user_id"`
	FirstName   string    `json:"first_name" db:"first_name"`
	LastName    string    `json:"last_name" db:"last_name"`
	Email       string    `json:"email" db:"email"`
	DateOfBirth time.Time `json:"date_of_birth" db:"date_of_birth"`
}
