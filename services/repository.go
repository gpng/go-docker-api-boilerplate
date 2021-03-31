package services

import (
	"github.com/gpng/go-docker-api-boilerplate/repository/interfaces"
	"github.com/gpng/go-docker-api-boilerplate/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	interfaces.UserRepository
}

func NewPostgresRespository(db *sqlx.DB) interfaces.Repository {
	return newRepository(db)
}

func newRepository(db *sqlx.DB) *repository {
	return &repository{
		postgres.NewUserRepository(),
	}
}
