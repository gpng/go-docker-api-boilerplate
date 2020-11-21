package handlers

import (
	vr "github.com/gpng/go-docker-api-boilerplate/services/validator"
	"go.uber.org/zap"
)

// Handlers struct
type Handlers struct {
	logger    *zap.Logger
	validator *vr.Validator
}

// New service
func New(
	logger *zap.Logger,
	validator *vr.Validator,
) *Handlers {
	return &Handlers{logger, validator}
}
