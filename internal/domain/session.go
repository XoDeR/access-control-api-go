package domain

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID               uuid.UUID  `json:"id"`
	UserID           uuid.UUID  `json:"user_id"`
	RefreshTokenHash string     `json:"-"`
	UserAgent        string     `json:"user_agent,omitempty"`
	IPAddress        string     `json:"ip_address,omitempty"`
	RevokedAt        *time.Time `json:"revoked_at,omitempty"`
	ExpiresAt        time.Time  `json:"expires_at"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}
