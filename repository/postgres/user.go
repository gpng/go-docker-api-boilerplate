package postgres

import (
	"github.com/gpng/go-docker-api-boilerplate/repository/interfaces"
	"github.com/gpng/go-docker-api-boilerplate/repository/models"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
}

func NewUserRepository() interfaces.UserRepository {
	return userRepository{}
}

func (r userRepository) AllUsers(db *sqlx.DB) ([]models.User, error) {
	users := make([]models.User, 0)

	err := db.Select(
		&users,
		`
			SELECT *
			  FROM users
		`,
	)

	return users, err
}
