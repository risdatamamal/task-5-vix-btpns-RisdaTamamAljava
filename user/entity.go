package user

import "time"

type User struct {
	ID              int
	Username        string
	Name            string
	Email           string
	// EmailVerifiedAt time.Time
	Password        string
	Avatar          string
	PhoneNumber     string
	City            string
	Address         string
	Roles           string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
