package main

import (
	"database/sql"
	"log"

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

	primaryKey, _ := uuid.NewUUID()
	user := model.User{}

	insertUserStmt.Exec()

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
