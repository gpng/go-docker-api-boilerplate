package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"go.uber.org/zap"
)

func (h *Handlers) optionalAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r, _ = h.validateToken(r)
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func (h *Handlers) requireAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r, err := h.validateToken(r)
		if err != nil {
			h.logger.Error("invalid token", zap.Error(err))
			respondWithStatus(w, http.StatusForbidden, message(err.Error()))
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func (h *Handlers) validateToken(r *http.Request) (*http.Request, error) {
	token := tokenFromHeader(r)
	if token == "" {
		return r, fmt.Errorf("missing token")
	}

	claims, err := h.jwt.ValidateAccessToken(token)
	if err != nil {
		h.logger.Error("invalid token", zap.Error(err))
		return r, fmt.Errorf("invalid token")
	}

	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		h.logger.Error("invalid user id", zap.String("userID", claims.UserID), zap.Error(err))
		return r, fmt.Errorf("invalid user id")
	}

	session, err := h.repo.GetSessionByUserID(h.db, userID)
	if err != nil {
		h.logger.Error("error retrieving session", zap.Error(err))
		return r, fmt.Errorf("database error")
	}

	if session.Invalidated {
		h.logger.Error("user banned", zap.String("userID", claims.UserID))
		return r, fmt.Errorf("user has been banned")
	}

	ctx := context.WithValue(r.Context(), contextKey(contextUID), claims.UserID)
	return r.WithContext(ctx), nil
}

// tokenFromHeader retrieves token string from Authorization Header
func tokenFromHeader(r *http.Request) string {
	bearer := r.Header.Get("Authorization")
	split := strings.SplitN(bearer, " ", 2)
	if len(split) > 1 || strings.ToUpper(split[0]) == "BEARER" {
		return split[1]
	}
	return ""
}
