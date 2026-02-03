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
	Role			Role `gorm:"default:ROLE_USER"`
	// TODO: add role field with ser/deser
	*Base
}

func (u User) TableName() string {
	return "users"
}