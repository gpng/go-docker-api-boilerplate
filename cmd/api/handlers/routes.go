package handlers

import (
	"github.com/go-chi/chi"
	_ "github.com/gpng/go-docker-api-boilerplate/docs" // required for generating docs
)

// Routes for app
func (h *Handlers) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", h.handleStatus())

	router.Mount("/auth", h.authRoutes())

	return router
}

func (h *Handlers) authRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/signup", h.handleSignup())
	r.Post("/login", h.handleLogin())

	return r
}
