package model

import "time"

// User ...
type User struct {
	id              string
	created         time.Time
	lastLogin       time.Time
	passwordChanged time.Time
}
