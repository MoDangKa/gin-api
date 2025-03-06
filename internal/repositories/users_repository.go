package repositories

import (
	"context"
	"errors"
	"fmt"
	"gin-api/internal/models"
	"gin-api/pkg/utils"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	query := `SELECT id, email, password, name, photo, role, active, created_at, updated_at FROM users`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var userList []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.Photo, &user.Role, &user.Active, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		userList = append(userList, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return userList, nil
}

func (r *UserRepository) CreateUser(user *models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = hashedPassword

	query := `
        INSERT INTO users (email, password, name, photo, role)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, active, created_at, updated_at
    `
	err = r.db.QueryRow(
		context.Background(),
		query,
		user.Email,
		user.Password,
		user.Name,
		user.Photo,
		user.Role,
	).Scan(
		&user.ID,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	query := `SELECT id, email, password, name, photo, role, created_at, updated_at FROM users WHERE id = $1`
	var user models.User
	err := r.db.QueryRow(context.Background(), query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.Photo,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	query := `
		UPDATE users
		SET password = $1, name = $2, photo = $3, role = $4, updated_at = $5
		WHERE id = $6, 
	`
	_, err := r.db.Exec(
		context.Background(),
		query,
		user.Password,
		user.Name,
		user.Photo,
		user.Role,
		time.Now(),
		user.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (r *UserRepository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func (r *UserRepository) LogIn(email string, password string) (*models.UserWithToken, error) {
	query := `SELECT id, email, password, name, photo, role, active, created_at, updated_at FROM users WHERE email = $1 AND active = true`
	var user models.User
	err := r.db.QueryRow(context.Background(), query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.Photo,
		&user.Role,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return nil, fmt.Errorf("incorrect email or password")
	}

	if !utils.CheckPassword(password, user.Password) {
		return nil, errors.New("incorrect email or password")
	}

	claims := jwt.MapClaims{
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Printf("Error signing token: %v", err)

		return nil, fmt.Errorf("could not generate token")
	}

	result := models.UserWithToken{
		User:  user,
		Token: tokenString,
	}

	return &result, nil
}
