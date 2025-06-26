package repositories

import (
	"context"
	"database/sql"
	"errors"
	"online-school/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type LessonRepository struct {
	db *sqlx.DB
}

func NewLessonRepository(db *sqlx.DB) *LessonRepository {
	return &LessonRepository{db: db}
}

func (r *LessonRepository) CreateLesson(ctx context.Context, lesson *models.Lesson) error {
	query := `INSERT INTO lessons (title, description, teacher_id, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`
	return r.db.QueryRowxContext(
		ctx,
		query,
		lesson.Title,
		lesson.Description,
		lesson.TeacherID,
		time.Now(),
		time.Now(),
	).Scan(&lesson.ID)
}

func (r *LessonRepository) GetLessonByID(ctx context.Context, id int) (*models.Lesson, error) {
	lesson := &models.Lesson{}
	query := `SELECT id, title, description, teacher_id, created_at, updated_at 
	          FROM lessons WHERE id = $1`
	err := r.db.GetContext(ctx, lesson, query, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("lesson not found")
	}
	return lesson, err
}

func (r *LessonRepository) GetLessonsByTeacher(ctx context.Context, teacherID int) ([]models.Lesson, error) {
	var lessons []models.Lesson
	query := `SELECT id, title, description, teacher_id, created_at, updated_at 
	          FROM lessons WHERE teacher_id = $1`
	err := r.db.SelectContext(ctx, &lessons, query, teacherID)
	return lessons, err
}

func (r *LessonRepository) GetLessonsForGroup(ctx context.Context, group string) ([]models.Lesson, error) {
	var lessons []models.Lesson
	query := `SELECT l.id, l.title, l.description, l.teacher_id 
	          FROM lessons l
	          JOIN schedules s ON l.id = s.lesson_id
	          WHERE s.group = $1`
	err := r.db.SelectContext(ctx, &lessons, query, group)
	return lessons, err
}
