package repositories

import (
	"context"
	"database/sql"
	"errors"
	"online-school/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type TeacherRepository struct {
	db *sqlx.DB
}

func NewTeacherRepository(db *sqlx.DB) *TeacherRepository {
	return &TeacherRepository{db: db}
}

func (r *TeacherRepository) CreateTeacher(ctx context.Context, teacher *models.Teacher) error {
	query := `INSERT INTO teachers (user_id, full_name, subject, approved, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	return r.db.QueryRowxContext(
		ctx,
		query,
		teacher.UserID,
		teacher.FullName,
		teacher.Subject,
		teacher.Approved,
		time.Now(),
		time.Now(),
	).Scan(&teacher.ID)
}

func (r *TeacherRepository) GetTeacherByID(ctx context.Context, id int) (*models.Teacher, error) {
	teacher := &models.Teacher{}
	query := `SELECT id, user_id, full_name, subject, approved, created_at, updated_at 
	          FROM teachers WHERE id = $1`
	err := r.db.GetContext(ctx, teacher, query, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("teacher not found")
	}
	return teacher, err
}

func (r *TeacherRepository) GetTeacherByUserID(ctx context.Context, userID int) (*models.Teacher, error) {
	teacher := &models.Teacher{}
	query := `SELECT id, user_id, full_name, subject, approved, created_at, updated_at 
	          FROM teachers WHERE user_id = $1`
	err := r.db.GetContext(ctx, teacher, query, userID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("teacher not found")
	}
	return teacher, err
}

func (r *TeacherRepository) ApproveTeacher(ctx context.Context, id int) error {
	query := `UPDATE teachers SET approved = true, updated_at = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, time.Now(), id)
	return err
}

func (r *TeacherRepository) GetTeachersBySubject(ctx context.Context, subject string) ([]models.Teacher, error) {
	var teachers []models.Teacher
	query := `SELECT id, user_id, full_name, subject, approved 
	          FROM teachers WHERE subject = $1 AND approved = true`
	err := r.db.SelectContext(ctx, &teachers, query, subject)
	return teachers, err
}
