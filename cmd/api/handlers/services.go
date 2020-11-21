package handlers

import (
	vr "github.com/gpng/go-docker-api-boilerplate/services/validator"
	"github.com/gpng/go-docker-api-boilerplate/sqlc/models"
	"go.uber.org/zap"
)

// Handlers struct
type Handlers struct {
	logger    *zap.Logger
	validator *vr.Validator
	db        models.DBTX
	repo      models.Querier
}

// New service
func New(
	logger *zap.Logger,
	validator *vr.Validator,
	db models.DBTX,
	repo models.Querier,
) *Handlers {
	return &Handlers{logger, validator, db, repo}
}
