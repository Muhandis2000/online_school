package repositories

import (
	"context"
	"database/sql"
	"errors"
	"online-school/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type StudentRepository struct {
	db *sqlx.DB
}

func NewStudentRepository(db *sqlx.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

func (r *StudentRepository) CreateStudent(ctx context.Context, student *models.Student) error {
	query := `INSERT INTO students (user_id, full_name, group_name, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`
	return r.db.QueryRowxContext(
		ctx,
		query,
		student.UserID,
		student.FullName,
		student.Group,
		time.Now(),
		time.Now(),
	).Scan(&student.ID)
}

func (r *StudentRepository) GetStudentByID(ctx context.Context, id int) (*models.Student, error) {
	student := &models.Student{}
	query := `SELECT id, user_id, full_name, group_name, created_at, updated_at 
	          FROM students WHERE id = $1`
	err := r.db.GetContext(ctx, student, query, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("student not found")
	}
	return student, err
}

func (r *StudentRepository) GetStudentByUserID(ctx context.Context, userID int) (*models.Student, error) {
	student := &models.Student{}
	query := `SELECT id, user_id, full_name, group_name, created_at, updated_at 
	          FROM students WHERE user_id = $1`
	err := r.db.GetContext(ctx, student, query, userID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("student not found")
	}
	return student, err
}

func (r *StudentRepository) GetStudentsByGroup(ctx context.Context, group string) ([]models.Student, error) {
	var students []models.Student
	query := `SELECT id, user_id, full_name, group_name 
	          FROM students WHERE group_name = $1`
	err := r.db.SelectContext(ctx, &students, query, group)
	return students, err
}
