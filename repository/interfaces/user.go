package interfaces

import (
	"github.com/gpng/go-docker-api-boilerplate/repository/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	AllUsers(*sqlx.DB) ([]models.User, error)
}
