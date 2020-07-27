package main

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/zmohling/wantaprice/pkg/model"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// NewService creates a identity service with necessary dependencies.
func NewService(db *sql.DB) IdentityService {
	return &identityService{
		DB: db,
	}
}

// IdentityService ...
type IdentityService interface {
	CreateUser(user model.User, password string) (id string, err error)
}

type identityService struct {
	DB *sql.DB
}

func (s *identityService) CreateUser(user model.User, password string) (id string, err error) {
	tx, err := s.DB.Begin()

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
	user.Password, _ = hashPassword(password)

	insertUserStmt.Exec(user.ID, user.Login, user.DisplayName, user.Phone, user.Created, user.LastLogin, user.PasswordChanged, user.Password)

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return user.ID, nil
}

func hashPassword(password string) (string, error) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), nil
}
