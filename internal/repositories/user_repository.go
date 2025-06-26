package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"online-school/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (username, email, password, role, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	return r.db.QueryRowxContext(
		ctx,
		query,
		user.Username,
		user.Email,
		user.Password,
		user.Role,
		time.Now(),
		time.Now(),
	).Scan(&user.ID)
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, email, password, role, created_at, updated_at 
	          FROM users WHERE email = $1`
	err := r.db.GetContext(ctx, user, query, email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("пользователь не найден")
	}
	return user, err
}
