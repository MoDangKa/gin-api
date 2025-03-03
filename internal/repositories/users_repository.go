package repositories

import (
	"context"
	"fmt"
	"gin-api/internal/models"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

// GetAllUsers fetches all users from the database
func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	query := `SELECT id, username, password, avatar, is_admin, created_at, updated_at FROM users`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var userList []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Avatar, &user.IsAdmin, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		userList = append(userList, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return userList, nil
}

// CreateUser inserts a new user into the database
func (r *UserRepository) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (username, password, avatar, is_admin, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	err := r.db.QueryRow(
		context.Background(),
		query,
		user.Username,
		user.Password,
		user.Avatar,
		user.IsAdmin,
		time.Now(),
		time.Now(),
	).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

// GetUserByID fetches a user by their ID
func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	query := `SELECT id, username, password, avatar, is_admin, created_at, updated_at FROM users WHERE id = $1`
	var user models.User
	err := r.db.QueryRow(context.Background(), query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Avatar,
		&user.IsAdmin,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

// UpdateUser updates a user's details
func (r *UserRepository) UpdateUser(user *models.User) error {
	query := `
		UPDATE users
		SET username = $1, password = $2, avatar = $3, is_admin = $4, updated_at = $5
		WHERE id = $6
	`
	_, err := r.db.Exec(
		context.Background(),
		query,
		user.Username,
		user.Password,
		user.Avatar,
		user.IsAdmin,
		time.Now(),
		user.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

// DeleteUser deletes a user by their ID
func (r *UserRepository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
