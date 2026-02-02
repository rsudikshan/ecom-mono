package types

import "time"

const (
	AUTH_USER_KEY = "AUTH_USER"
)

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

type AuthUser struct {
	
}