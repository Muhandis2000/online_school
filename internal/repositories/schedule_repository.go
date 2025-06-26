package repositories

import (
	"context"
	"online-school/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type ScheduleRepository struct {
	db *sqlx.DB
}

func NewScheduleRepository(db *sqlx.DB) *ScheduleRepository {
	return &ScheduleRepository{db: db}
}
func (r *ScheduleRepository) CreateSchedule(ctx context.Context, schedule *models.Schedule) error {
	query := `INSERT INTO schedules (lesson_id, group_name, start_time, end_time, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	return r.db.QueryRowxContext(
		ctx,
		query,
		schedule.LessonID,
		schedule.Group,
		schedule.StartTime,
		schedule.EndTime,
		time.Now(),
		time.Now(),
	).Scan(&schedule.ID)
}

func (r *ScheduleRepository) GetScheduleByLesson(ctx context.Context, lessonID int) ([]models.Schedule, error) {
	var schedules []models.Schedule
	query := `SELECT id, lesson_id, group_name, start_time, end_time 
	          FROM schedules WHERE lesson_id = $1`
	err := r.db.SelectContext(ctx, &schedules, query, lessonID)
	return schedules, err
}

func (r *ScheduleRepository) GetScheduleForGroup(ctx context.Context, group string) ([]models.Schedule, error) {
	var schedules []models.Schedule
	query := `SELECT id, lesson_id, group_name, start_time, end_time 
	          FROM schedules WHERE group_name = $1`
	err := r.db.SelectContext(ctx, &schedules, query, group)
	return schedules, err
}

// В ScheduleRepository
func (r *ScheduleRepository) GetScheduleByTeacher(ctx context.Context, teacherID int) ([]models.Schedule, error) {
	query := `
		SELECT s.id, s.lesson_id, s.group_name, s.start_time, s.end_time, s.created_at, s.updated_at
		FROM schedules s
		JOIN lessons l ON s.lesson_id = l.id
		WHERE l.teacher_id = $1
	`

	var schedules []models.Schedule
	err := r.db.SelectContext(ctx, &schedules, query, teacherID)
	return schedules, err
}
