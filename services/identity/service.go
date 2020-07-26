package main

import "github.com/zmohling/wantaprice/pkg/model"

type Service interface {
	CreateUser(user model.User)
}
