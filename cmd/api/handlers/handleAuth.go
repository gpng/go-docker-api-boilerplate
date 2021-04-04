package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gpng/go-docker-api-boilerplate/repository/models"
	"go.uber.org/zap"
)

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Response struct {
	Tokens Tokens      `json:"tokens"`
	User   models.User `json:"user"`
}

func (h *Handlers) handleSignup() http.HandlerFunc {
	type request struct {
		Email    string `json:"email" validate:"email"`
		Password string `json:"password" validate:"required"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		params := request{}
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			h.logger.Error("error decoding request body", zap.Error(err))
			respondWithStatus(w, http.StatusBadRequest, message("unable to decode request body"))
			return
		}

		err = h.validator.Validate(params)
		if err != nil {
			respondWithStatus(w, http.StatusBadRequest, message(h.validator.TranslateValidatorErr(err)))
			return
		}

		// check if email used
		_, err = h.repo.GetUserByEmail(h.db, params.Email)
		if err == nil {
			respondWithStatus(w, http.StatusBadRequest, message("email already in use"))
			return
		}
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			h.logger.Error("error getting user by email", zap.Error(err))
			respondWithStatus(w, http.StatusInternalServerError, message("database error"))
			return
		}

		hashedPassword, err := hashPassword(params.Password)
		if err != nil {
			h.logger.Error("error hashing password", zap.Error(err))
			respondWithStatus(w, http.StatusInternalServerError, message("invalid password"))
			return
		}

		user, err := h.repo.CreateUser(h.db, params.Email, hashedPassword)
		if err != nil {
			h.logger.Error("error creating user", zap.Error(err))
			respondWithStatus(w, http.StatusInternalServerError, message("database error"))
			return
		}

		// create tokens
		tokens, err := h.generateTokens(user.ID)
		if err != nil {
			h.logger.Error("error generating tokens", zap.Error(err))
			respondWithStatus(w, http.StatusInternalServerError, message("tokens error"))
			return
		}

		respond(w, dataMessage(Response{
			Tokens: tokens,
			User:   user,
		}, "success"))
	}
}

func (h *Handlers) handleLogin() http.HandlerFunc {
	type request struct {
		Email    string `json:"email" validate:"email"`
		Password string `json:"password" validate:"required"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		params := request{}
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			h.logger.Error("error decoding request body", zap.Error(err))
			respondWithStatus(w, http.StatusBadRequest, message("unable to decode request body"))
			return
		}

		err = h.validator.Validate(params)
		if err != nil {
			respondWithStatus(w, http.StatusBadRequest, message(h.validator.TranslateValidatorErr(err)))
			return
		}

		// check if email used
		user, err := h.repo.GetUserByEmail(h.db, params.Email)
		if errors.Is(err, sql.ErrNoRows) {
			respondWithStatus(w, http.StatusBadRequest, message("invalid password"))
			return
		}
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			h.logger.Error("error getting user by email", zap.Error(err))
			respondWithStatus(w, http.StatusInternalServerError, message("database error"))
			return
		}

		err = checkPassword(user.Password, params.Password)
		if err != nil {
			respondWithStatus(w, http.StatusInternalServerError, message("invalid password"))
			return
		}

		// create tokens
		tokens, err := h.generateTokens(user.ID)
		if err != nil {
			h.logger.Error("error generating tokens", zap.Error(err))
			respondWithStatus(w, http.StatusInternalServerError, message("tokens error"))
			return
		}

		respond(w, dataMessage(Response{
			Tokens: tokens,
			User:   user,
		}, "success"))
	}
}

func (h *Handlers) handleRefresh() http.HandlerFunc {
	type request struct {
		Token string `json:"token" validate:"required"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		params := request{}
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			h.logger.Error("error decoding request body", zap.Error(err))
			respondWithStatus(w, http.StatusBadRequest, message("unable to decode request body"))
			return
		}

		err = h.validator.Validate(params)
		if err != nil {
			respondWithStatus(w, http.StatusBadRequest, message(h.validator.TranslateValidatorErr(err)))
			return
		}

		// validate refresh token
		claims, err := h.jwt.ValidateRefreshToken(params.Token)
		if err != nil {
			h.logger.Error("invalid token", zap.Error(err))
			respondWithStatus(w, http.StatusBadRequest, message(err.Error()))
			return
		}

		userID, err := uuid.Parse(claims.UserID)
		if err != nil {
			h.logger.Error("invalid user ID", zap.Error(err), zap.String("userID", claims.UserID))
			respondWithStatus(w, http.StatusBadRequest, message("invalid user id"))
			return
		}

		// check session
		session, err := h.repo.GetSessionByUserID(h.db, userID)
		if err != nil {
			h.logger.Error("error retrieving session", zap.Error(err))
			respondWithStatus(w, http.StatusForbidden, message("database error"))
			return
		}

		if session.Invalidated {
			respondWithStatus(w, http.StatusForbidden, message("user has been banned"))
			return
		}
		if session.RefreshToken != params.Token {
			respondWithStatus(w, http.StatusForbidden, message("invalid or expired token"))
			return
		}

		user, err := h.repo.GetUserByID(h.db, userID)
		if err != nil {
			h.logger.Error("error getting user", zap.Error(err))
			respondWithStatus(w, http.StatusInternalServerError, message("database error"))
			return
		}

		// create tokens
		tokens, err := h.generateTokens(user.ID)
		if err != nil {
			h.logger.Error("error generating tokens", zap.Error(err))
			respondWithStatus(w, http.StatusInternalServerError, message("tokens error"))
			return
		}

		respond(w, dataMessage(Response{
			Tokens: tokens,
			User:   user,
		}, "success"))
	}
}

func (h *Handlers) generateTokens(userID uuid.UUID) (Tokens, error) {
	tokens := Tokens{}
	accessToken, refreshToken, err := h.jwt.GenerateTokenPair(userID)
	if err != nil {
		return tokens, err
	}

	session, err := h.repo.CreateSession(h.db, userID, refreshToken)
	if err != nil {
		return tokens, err
	}

	if session.Invalidated {
		return tokens, ErrTokenInvalidated
	}

	tokens.AccessToken = accessToken
	tokens.RefreshToken = refreshToken

	return tokens, err
}
