package model

import "time"

// User ...
type User struct {
	ID              string
	Login           string
	DisplayName     string
	Phone           string
	Created         time.Time
	LastLogin       time.Time
	PasswordChanged time.Time
	salt            string
	password        string
}
