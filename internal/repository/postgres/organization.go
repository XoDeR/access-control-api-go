package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/XoDeR/access-control-api-go/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (db *DB) CreateOrganization(ctx context.Context, org *domain.Organization) error {
	query := `INSERT INTO organizations (id, name, slug, owner_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())`
	_, err := db.Pool.Exec(ctx, query, org.ID, org.Name, org.Slug, org.OwnerID)
	if err != nil {
		return fmt.Errorf("create organization: %w", err)
	}
	return nil
}

func (db *DB) GetOrganizationByID(ctx context.Context, id uuid.UUID) (*domain.Organization, error) {
	query := `SELECT id, name, slug, owner_id, created_at, updated_at, deleted_at
		FROM organizations WHERE id = $1 AND deleted_at IS NULL`
	var o domain.Organization
	err := db.Pool.QueryRow(ctx, query, id).Scan(
		&o.ID, &o.Name, &o.Slug, &o.OwnerID, &o.CreatedAt, &o.UpdatedAt, &o.DeletedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("get organization: %w", err)
	}
	return &o, nil
}

func (db *DB) SoftDeleteOrganization(ctx context.Context, orgID uuid.UUID) error {
	query := `UPDATE organizations SET deleted_at = NOW(), updated_at = NOW() WHERE id = $1 AND deleted_at IS NULL`
	tag, err := db.Pool.Exec(ctx, query, orgID)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
