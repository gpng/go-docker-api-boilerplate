package handlers

import (
	"github.com/go-chi/chi"
	_ "github.com/gpng/go-docker-api-boilerplate/docs" // required for generating docs
)

// Routes for app
func (s *Handlers) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", s.handleStatus())

	router.Get("/users", s.handleUsers())

	return router
}
