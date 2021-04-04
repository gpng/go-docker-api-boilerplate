package interfaces

import (
	"github.com/google/uuid"
	"github.com/gpng/go-docker-api-boilerplate/repository/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(db *sqlx.DB, email, password string) (models.User, error)
	GetUserByID(db *sqlx.DB, id uuid.UUID) (models.User, error)
	GetUserByEmail(db *sqlx.DB, email string) (models.User, error)
}
