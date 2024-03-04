package models

type User struct {
	ID             int
	Username       string
	HashedPassword string
	Email          string
	Phone          string
	IsAdmin        bool
}
