package user

import "time"

type User struct {
	ID              int
	Username        string
	Name            string
	Email           string
	Password        string
	Avatar          string
	PhoneNumber     string
	City            string
	Address         string
	Roles           string
	EmailVerifiedAt time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
