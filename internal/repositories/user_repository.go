package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"online-school/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepositoryInterface interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
}

type TeacherRepositoryInterface interface {
	CreateTeacher(teacher *models.Teacher) error
	ApproveTeacher(id int) error
	GetTeachers() ([]models.Teacher, error)
}

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
func (r *UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	query := `SELECT id, username, email, password, role, created_at, updated_at FROM users`
	if err := r.db.SelectContext(ctx, &users, query); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, email, password, role, created_at, updated_at FROM users WHERE id = $1`
	if err := r.db.GetContext(ctx, &user, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("пользователь не найден")
		}
		return nil, err
	}
	return &user, nil
}
