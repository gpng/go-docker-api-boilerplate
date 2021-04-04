package handlers

import (
	"github.com/gpng/go-docker-api-boilerplate/repository/interfaces"
	"github.com/gpng/go-docker-api-boilerplate/services/jwt"
	vr "github.com/gpng/go-docker-api-boilerplate/services/validator"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Handlers struct
type Handlers struct {
	logger    *zap.Logger
	validator *vr.Validator
	jwt       *jwt.Jwt
	db        *sqlx.DB
	repo      interfaces.Repository
}

// New service
func New(
	logger *zap.Logger,
	validator *vr.Validator,
	jwt *jwt.Jwt,
	db *sqlx.DB,
	repo interfaces.Repository,
) *Handlers {
	return &Handlers{logger, validator, jwt, db, repo}
}
