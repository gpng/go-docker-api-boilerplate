package handlers

import (
	"github.com/gpng/go-docker-api-boilerplate/repository/interfaces"
	vr "github.com/gpng/go-docker-api-boilerplate/services/validator"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Handlers struct
type Handlers struct {
	logger    *zap.Logger
	validator *vr.Validator
	db        *sqlx.DB
	repo      interfaces.Repository
}

// New service
func New(
	logger *zap.Logger,
	validator *vr.Validator,
	db *sqlx.DB,
	repo interfaces.Repository,
) *Handlers {
	return &Handlers{logger, validator, db, repo}
}
