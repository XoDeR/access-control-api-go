package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/XoDeR/access-control-api-go/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (db *DB) CreateUser(ctx context.Context, u *domain.User) error {
	query := `INSERT INTO users (id, email, password_hash, name, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())`
	_, err := db.Pool.Exec(ctx, query, u.ID, u.Email, u.PasswordHash, u.Name)
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}
	return nil
}

func (db *DB) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT id, email, password_hash, name, created_at, updated_at, deleted_at
		FROM users WHERE email = $1 AND deleted_at IS NULL`
	var u domain.User
	err := db.Pool.QueryRow(ctx, query, email).Scan(
		&u.ID, &u.Email, &u.PasswordHash, &u.Name, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("get user by email: %w", err)
	}
	return &u, nil
}

func (db *DB) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	query := `SELECT id, email, password_hash, name, created_at, updated_at, deleted_at
		FROM users WHERE id = $1 AND deleted_at IS NULL`
	var u domain.User
	err := db.Pool.QueryRow(ctx, query, id).Scan(
		&u.ID, &u.Email, &u.PasswordHash, &u.Name, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("get user by id: %w", err)
	}
	return &u, nil
}

func (db *DB) UpdateUserPassword(ctx context.Context, userID uuid.UUID, hash string) error {
	query := `UPDATE users SET password_hash = $2, updated_at = NOW() WHERE id = $1 AND deleted_at IS NULL`
	_, err := db.Pool.Exec(ctx, query, userID, hash)
	return err
}

func (db *DB) SoftDeleteUser(ctx context.Context, userID uuid.UUID) error {
	query := `UPDATE users SET deleted_at = NOW(), updated_at = NOW() WHERE id = $1 AND deleted_at IS NULL`
	tag, err := db.Pool.Exec(ctx, query, userID)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
