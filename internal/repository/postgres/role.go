package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/XoDeR/access-control-api-go/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (db *DB) GetRoleByName(ctx context.Context, name string) (*domain.Role, error) {
	query := `SELECT id, name, created_at FROM roles WHERE name = $1`
	var r domain.Role
	err := db.Pool.QueryRow(ctx, query, name).Scan(&r.ID, &r.Name, &r.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("get role: %w", err)
	}
	return &r, nil
}

func (db *DB) GetRoleByID(ctx context.Context, id uuid.UUID) (*domain.Role, error) {
	query := `SELECT id, name, created_at FROM roles WHERE id = $1`
	var r domain.Role
	err := db.Pool.QueryRow(ctx, query, id).Scan(&r.ID, &r.Name, &r.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("get role by id: %w", err)
	}
	return &r, nil
}

func (db *DB) ListRoles(ctx context.Context) ([]domain.Role, error) {
	rows, err := db.Pool.Query(ctx, `SELECT id, name, created_at FROM roles ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var roles []domain.Role
	for rows.Next() {
		var r domain.Role
		if err := rows.Scan(&r.ID, &r.Name, &r.CreatedAt); err != nil {
			return nil, err
		}
		roles = append(roles, r)
	}
	return roles, rows.Err()
}
