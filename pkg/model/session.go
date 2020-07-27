package model

import "time"

// Session ...
type Session struct {
	ID        string
	Login     string
	UserID    string
	ExpiresAt time.Time
}
