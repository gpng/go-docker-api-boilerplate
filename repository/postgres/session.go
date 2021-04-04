package postgres

import (
	"github.com/google/uuid"
	"github.com/gpng/go-docker-api-boilerplate/repository/interfaces"
	"github.com/gpng/go-docker-api-boilerplate/repository/models"
	"github.com/jmoiron/sqlx"
)

type sessionRepository struct {
}

func NewSessionRepository() interfaces.SessionRepository {
	return sessionRepository{}
}

func (r sessionRepository) CreateSession(db *sqlx.DB, userID uuid.UUID, refreshToken string) (models.Session, error) {
	session := models.Session{}

	err := db.Get(
		&session,
		`
			INSERT INTO sessions (user_id, refresh_token)
			VALUES ($1, $2)
      ON CONFLICT (user_id)
				DO UPDATE
							SET refresh_token = $2 
			RETURNING *
		`, userID.String(), refreshToken,
	)

	return session, err
}
