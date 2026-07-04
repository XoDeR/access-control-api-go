package domain

import (
	"time"

	"github.com/google/uuid"
)

type APIKey struct {
	ID             uuid.UUID  `json:"id"`
	OrganizationID uuid.UUID  `json:"organization_id"`
	Name           string     `json:"name"`
	KeyPrefix      string     `json:"key_prefix"`
	KeyHash        string     `json:"-"`
	CreatedBy      uuid.UUID  `json:"created_by"`
	LastUsedAt     *time.Time `json:"last_used_at,omitempty"`
	RevokedAt      *time.Time `json:"revoked_at,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
}

type APIKeyCreate struct {
	Name string `json:"name" validate:"required,min=1"`
}
