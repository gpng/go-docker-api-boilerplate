package handlers

import (
	"net/http"

	"go.uber.org/zap"
)

// handleStatus returns the current api version
func (s *Handlers) handleUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := s.repo.GetUsers(r.Context())
		if err != nil {
			s.logger.Error("error getting users", zap.Error(err))
			respondWithStatus(w, http.StatusInternalServerError, message("database error"))
			return
		}
		respond(w, dataMessage(users, "successfully retrieved users"))
	}
}
