package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) IsActive(email string) (bool, error) {
	query := `SELECT active FROM public.users WHERE email = $1`
	var active bool

	err := r.db.QueryRow(context.Background(), query, email).Scan(&active)
	if err != nil {
		return false, fmt.Errorf("error fetching user active status: %w", err)
	}

	return active, nil
}
