package users

import "time"

type RegistrationForm struct {
	Email    string
	Name     string
	Password string
}

type LoginForm struct {
	Email    string
	Password string
}

type User struct {
	Id         int
	Email      string
	Password   string
	Name       string
	Registered time.Time
}
