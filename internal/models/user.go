package models

type User struct {
	ID             int
	Username       string
	HashedPassword string
	Email          string
	Phone          string
}

func NewUser(username, hashedPassword, email, phone string) *User {
	return &User{
		Username:       username,
		HashedPassword: hashedPassword,
		Email:          email,
		Phone:          phone,
	}
}
