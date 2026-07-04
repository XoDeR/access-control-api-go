package domain

import (
	"time"

	"github.com/google/uuid"
)

type Invitation struct {
	ID             uuid.UUID  `json:"id"`
	OrganizationID uuid.UUID  `json:"organization_id"`
	Email          string     `json:"email"`
	RoleID         uuid.UUID  `json:"role_id"`
	RoleName       string     `json:"role_name,omitempty"`
	TokenHash      string     `json:"-"`
	InvitedBy      uuid.UUID  `json:"invited_by"`
	ExpiresAt      time.Time  `json:"expires_at"`
	AcceptedAt     *time.Time `json:"accepted_at,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
}

type InvitationCreate struct {
	Email    string `json:"email" validate:"required,email"`
	RoleName string `json:"role" validate:"required,oneof=admin member viewer"`
}
