package domain

import (
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Slug      string     `json:"slug"`
	OwnerID   uuid.UUID  `json:"owner_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type OrganizationCreate struct {
	Name string `json:"name" validate:"required,min=1"`
	Slug string `json:"slug" validate:"required,min=1"`
}
