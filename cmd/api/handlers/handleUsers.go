package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"go.uber.org/zap"
)

// handleStatus returns the current api version
func (h *Handlers) handleUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := h.repo.AllUsers(h.db)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			h.logger.Error("error getting all users", zap.Error(err))
			respondWithStatus(w, http.StatusInternalServerError, message("database error"))
			return
		}
		respond(w, dataMessage(users, "success"))
	}
}
