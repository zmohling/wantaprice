package model

import "time"

// User ...
type User struct {
	ID              string
	Created         time.Time
	LastLogin       time.Time
	PasswordChanged time.Time
}
