package model

import "time"

// Session ...
type Session struct {
	id        string
	login     string
	userID    string
	expiresAt time.Time
}
