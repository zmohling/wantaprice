package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/zmohling/wantaprice/pkg/model"
)

// IdentityService ...
type IdentityService interface {
	CreateUser(user model.User, password string) (model.User, error)
}

type identityService struct {
	db sql.DB
}

func (s *identityService) CreateUser(user model.User, password string) (model.User, error) {
	tx, err := s.db.Begin()

	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	insertUserStmt, err := tx.Prepare("INSERT INTO `users` VALUES (1, 2, 3, 4)")
	if err != nil {
		log.Fatal(err)
	}
	defer insertUserStmt.Close()

	pk, _ := uuid.NewUUID()
	user.ID = pk.String()
	user.Created = time.Now()
	user.PasswordChanged = time.Now()

	insertUserStmt.Exec(user.ID, user.Created, user.LastLogin, user.PasswordChanged)

	insertProfile, err := tx.Prepare("INSERT INTO `profiles` VALUES (1, 2, 3, 4)")
	if err != nil {
		log.Fatal(err)
	}
	defer insertProfile.Close()

	insertProfile.Exec(profile.L)

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
