package interfaces

import (
	"github.com/google/uuid"
	"github.com/gpng/go-docker-api-boilerplate/repository/models"
	"github.com/jmoiron/sqlx"
)

type SessionRepository interface {
	CreateSession(db *sqlx.DB, userID uuid.UUID, refreshToken string) (models.Session, error)
}
