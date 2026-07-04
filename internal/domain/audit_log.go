package domain

import (
	"time"

	"github.com/google/uuid"
)

type AuditAction string

const (
	AuditLogin             AuditAction = "login"
	AuditInvite            AuditAction = "invite"
	AuditRoleChange        AuditAction = "role_change"
	AuditPermissionChange  AuditAction = "permission_change"
	AuditLogout            AuditAction = "logout"
	AuditPasswordReset     AuditAction = "password_reset"
	AuditAPIKeyCreate      AuditAction = "api_key_create"
	AuditAPIKeyDelete      AuditAction = "api_key_delete"
	AuditOrgDelete         AuditAction = "org_delete"
	AuditUserDelete        AuditAction = "user_delete"
)

type AuditLog struct {
	ID             uuid.UUID              `json:"id"`
	OrganizationID *uuid.UUID             `json:"organization_id,omitempty"`
	UserID         uuid.UUID              `json:"user_id"`
	Action         AuditAction            `json:"action"`
	ResourceType   string                 `json:"resource_type,omitempty"`
	ResourceID     *uuid.UUID             `json:"resource_id,omitempty"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
	IPAddress      string                 `json:"ip_address,omitempty"`
	CreatedAt      time.Time              `json:"created_at"`
}
