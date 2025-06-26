package repositories

import (
	"context"
	"online-school/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type GradeRepository struct {
	db *sqlx.DB
}

func NewGradeRepository(db *sqlx.DB) *GradeRepository {
	return &GradeRepository{db: db}
}

func (r *GradeRepository) CreateGrade(ctx context.Context, grade *models.Grade) error {
	query := `INSERT INTO grades (student_id, lesson_id, value, description, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	return r.db.QueryRowxContext(
		ctx,
		query,
		grade.StudentID,
		grade.LessonID,
		grade.Value,
		grade.Description,
		time.Now(),
		time.Now(),
	).Scan(&grade.ID)
}

func (r *GradeRepository) GetGradesByStudent(ctx context.Context, studentID int) ([]models.Grade, error) {
	var grades []models.Grade
	query := `SELECT id, student_id, lesson_id, value, description 
	          FROM grades WHERE student_id = $1`
	err := r.db.SelectContext(ctx, &grades, query, studentID)
	return grades, err
}
