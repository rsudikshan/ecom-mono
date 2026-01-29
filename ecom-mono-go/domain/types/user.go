package types

import "time"

type User struct {
	ID 				ID
	Email           string
	Username        string
	Password        string
	EmailVerified   bool
	PasswordResetAt *time.Time

	Base
}

func (u User) TableName() string {
	return "users"
}