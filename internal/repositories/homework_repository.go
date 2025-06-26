package repositories

import (
	"context"
	"online-school/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type HomeworkRepository struct {
	db *sqlx.DB
}

func NewHomeworkRepository(db *sqlx.DB) *HomeworkRepository {
	return &HomeworkRepository{db: db}
}

func (r *HomeworkRepository) CreateHomework(ctx context.Context, homework *models.Homework) error {
	query := `INSERT INTO homeworks (lesson_id, student_id, description, submitted_at, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	return r.db.QueryRowxContext(
		ctx,
		query,
		homework.LessonID,
		homework.StudentID,
		homework.Description,
		homework.SubmittedAt,
		time.Now(),
		time.Now(),
	).Scan(&homework.ID)
}

func (r *HomeworkRepository) GetHomeworkByID(ctx context.Context, id int) (*models.Homework, error) {
	homework := &models.Homework{}
	query := `SELECT id, lesson_id, student_id, description, submitted_at 
	          FROM homeworks WHERE id = $1`
	err := r.db.GetContext(ctx, homework, query, id)
	return homework, err
}
