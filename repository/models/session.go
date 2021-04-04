package models

import "github.com/google/uuid"

type Session struct {
	UserID       uuid.UUID `db:"user_id"`
	RefreshToken string    `db:"refresh_token"`
	Invalidated  bool
}
