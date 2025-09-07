package user

import "time"

type UserDetails struct {
	ID          int64     `json:"id" db:"id"`
	UserId      int64     `json:"userId" db:"user_id"`
	FirstName   string    `json:"firstName" db:"first_name"`
	LastName    string    `json:"lastName" db:"last_name"`
	Email       string    `json:"email" db:"email"`
	DateOfBirth time.Time `json:"dateOfBirth" db:"date_of_birth"`
}
