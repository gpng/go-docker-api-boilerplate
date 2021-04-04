package postgres

import (
	"github.com/google/uuid"
	"github.com/gpng/go-docker-api-boilerplate/repository/interfaces"
	"github.com/gpng/go-docker-api-boilerplate/repository/models"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
}

func NewUserRepository() interfaces.UserRepository {
	return userRepository{}
}

func (r userRepository) GetUserByID(db *sqlx.DB, id uuid.UUID) (models.User, error) {
	user := models.User{}

	err := db.Get(
		&user,
		`
			SELECT *
			FROM users
			WHERE id = $1
		`, id.String(),
	)

	return user, err
}

func (r userRepository) GetUserByEmail(db *sqlx.DB, email string) (models.User, error) {
	user := models.User{}

	err := db.Get(
		&user,
		`
			SELECT *
			FROM users
			WHERE email = $1
		`, email,
	)

	return user, err
}

func (r userRepository) CreateUser(db *sqlx.DB, email, password string) (models.User, error) {
	user := models.User{}

	err := db.Get(
		&user,
		`
			INSERT INTO users (email, password)
			VALUES ($1, $2)
			RETURNING *
		`, email, password,
	)

	return user, err
}
