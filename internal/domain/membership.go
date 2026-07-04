package domain

import (
	"time"

	"github.com/google/uuid"
)

type Membership struct {
	ID             uuid.UUID `json:"id"`
	OrganizationID uuid.UUID `json:"organization_id"`
	UserID         uuid.UUID `json:"user_id"`
	RoleID         uuid.UUID `json:"role_id"`
	RoleName       string    `json:"role_name,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type MembershipWithUser struct {
	Membership
	Email string `json:"email"`
	Name  string `json:"name"`
}
