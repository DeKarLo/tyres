package models

import (
	"time"
)

type User struct {
	ID             int
	Username       string
	HashedPassword string
	Email          string
	Phone          string
	IsAdmin        bool
	CreatedAt      time.Time
	UpdatedAt 	   time.Time
}

func NewUser(username, hashedPassword, email, phone string, isAdmin bool, createdAt, updatedAt time.Time) *User {
	return &User{
		Username:       username,
		HashedPassword: hashedPassword,
		Email:          email,
		Phone:          phone,
		IsAdmin:        isAdmin,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}
}
