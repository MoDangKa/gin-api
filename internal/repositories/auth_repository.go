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

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) LogIn(email string, password string) (*models.Auth, error) {
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

	result := models.Auth{
		User:  user,
		Token: tokenString,
	}

	return &result, nil
}
