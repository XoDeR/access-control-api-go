package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/XoDeR/access-control-api-go/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (db *DB) CreateMembership(ctx context.Context, m *domain.Membership) error {
	query := `INSERT INTO memberships (id, org_id, user_id, role_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())`
	_, err := db.Pool.Exec(ctx, query, m.ID, m.OrganizationID, m.UserID, m.RoleID)
	return err
}

func (db *DB) GetMembership(ctx context.Context, orgID, userID uuid.UUID) (*domain.Membership, error) {
	query := `SELECT m.id, m.org_id, m.user_id, m.role_id, r.name, m.created_at, m.updated_at
		FROM memberships m
		JOIN roles r ON r.id = m.role_id
		WHERE m.org_id = $1 AND m.user_id = $2`
	var m domain.Membership
	err := db.Pool.QueryRow(ctx, query, orgID, userID).Scan(
		&m.ID, &m.OrganizationID, &m.UserID, &m.RoleID, &m.RoleName, &m.CreatedAt, &m.UpdatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("get membership: %w", err)
	}
	return &m, nil
}

func (db *DB) UpdateMembershipRole(ctx context.Context, orgID, userID, roleID uuid.UUID) error {
	query := `UPDATE memberships SET role_id = $3, updated_at = NOW()
		WHERE org_id = $1 AND user_id = $2`
	tag, err := db.Pool.Exec(ctx, query, orgID, userID, roleID)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (db *DB) ListMembers(ctx context.Context, orgID uuid.UUID, p domain.PageParams) ([]domain.MembershipWithUser, int, error) {
	countQuery := `SELECT COUNT(*) FROM memberships m JOIN users u ON u.id = m.user_id
		WHERE m.org_id = $1 AND u.deleted_at IS NULL`
	var total int
	if err := db.Pool.QueryRow(ctx, countQuery, orgID).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `SELECT m.id, m.org_id, m.user_id, m.role_id, r.name, m.created_at, m.updated_at,
		u.email, u.name
		FROM memberships m
		JOIN roles r ON r.id = m.role_id
		JOIN users u ON u.id = m.user_id
		WHERE m.org_id = $1 AND u.deleted_at IS NULL
		ORDER BY m.created_at DESC
		LIMIT $2 OFFSET $3`
	rows, err := db.Pool.Query(ctx, query, orgID, p.Limit, domain.Offset(p))
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var members []domain.MembershipWithUser
	for rows.Next() {
		var m domain.MembershipWithUser
		if err := rows.Scan(
			&m.ID, &m.OrganizationID, &m.UserID, &m.RoleID, &m.RoleName,
			&m.CreatedAt, &m.UpdatedAt, &m.Email, &m.Name,
		); err != nil {
			return nil, 0, err
		}
		members = append(members, m)
	}
	return members, total, rows.Err()
}
