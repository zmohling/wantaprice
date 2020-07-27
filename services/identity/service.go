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
	CreateUser(profile model.Profile, password string) (model.User, error)
}

type identityService struct {
	db sql.DB
}

func (s *identityService) CreateUser(profile model.Profile, password string) (model.User, error) {
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
	user := model.User{
		ID:              pk.String(),
		Created:         time.Now(),
		PasswordChanged: time.Now(),
	}

	insertUserStmt.Exec(user.ID, user.Created, user.LastLogin, user.PasswordChanged)

	insertProfile, err := tx.Prepare("INSERT INTO `profiles` VALUES (1, 2, 3, 4)")
	if err != nil {
		log.Fatal(err)
	}
	defer insertProfile.Close()

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
