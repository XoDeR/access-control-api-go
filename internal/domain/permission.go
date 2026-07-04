package domain

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

const (
	PermUsersRead     = "users.read"
	PermUsersInvite   = "users.invite"
	PermProjectsWrite = "projects.write"
	PermBillingRead   = "billing.read"
)
