package model

import "time"

// User ...
type User struct {
	ID              string    `json:"id"`
	Login           string    `json:"login"`
	DisplayName     string    `json:"displayName"`
	Phone           string    `json:"phone,omitempty"`
	Created         time.Time `json:"created"`
	LastLogin       time.Time `json:"-"`
	PasswordChanged time.Time `json:"-"`
	Password        string    `json:"-"`
}
