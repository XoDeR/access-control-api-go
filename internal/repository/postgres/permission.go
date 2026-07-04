package postgres

import (
	"context"
	"fmt"

	"github.com/XoDeR/access-control-api-go/internal/domain"
	"github.com/google/uuid"
)

func (db *DB) GetPermissionByName(ctx context.Context, name string) (*domain.Permission, error) {
	query := `SELECT id, name, created_at FROM permissions WHERE name = $1`
	var p domain.Permission
	err := db.Pool.QueryRow(ctx, query, name).Scan(&p.ID, &p.Name, &p.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("get permission: %w", err)
	}
	return &p, nil
}

func (db *DB) UserHasPermission(ctx context.Context, orgID, userID uuid.UUID, permName string) (bool, error) {
	query := `SELECT EXISTS (
		SELECT 1 FROM memberships m
		JOIN role_permissions rp ON rp.role_id = m.role_id
		JOIN permissions p ON p.id = rp.permission_id
		WHERE m.organization_id = $1 AND m.user_id = $2 AND p.name = $3
	)`
	var exists bool
	err := db.Pool.QueryRow(ctx, query, orgID, userID, permName).Scan(&exists)
	return exists, err
}

func (db *DB) ListPermissionsForRole(ctx context.Context, roleID uuid.UUID) ([]domain.Permission, error) {
	query := `SELECT p.id, p.name, p.created_at
		FROM permissions p
		JOIN role_permissions rp ON rp.permission_id = p.id
		WHERE rp.role_id = $1`
	rows, err := db.Pool.Query(ctx, query, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var perms []domain.Permission
	for rows.Next() {
		var p domain.Permission
		if err := rows.Scan(&p.ID, &p.Name, &p.CreatedAt); err != nil {
			return nil, err
		}
		perms = append(perms, p)
	}
	return perms, rows.Err()
}
