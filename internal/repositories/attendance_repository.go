package repositories

import (
	"context"
	"online-school/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type AttendanceRepository struct {
	db *sqlx.DB
}

func NewAttendanceRepository(db *sqlx.DB) *AttendanceRepository {
	return &AttendanceRepository{db: db}
}

func (r *AttendanceRepository) CreateAttendance(ctx context.Context, attendance *models.Attendance) error {
	query := `INSERT INTO attendances (student_id, lesson_id, present, date, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	return r.db.QueryRowxContext(
		ctx,
		query,
		attendance.StudentID,
		attendance.LessonID,
		attendance.Present,
		attendance.Date,
		time.Now(),
		time.Now(),
	).Scan(&attendance.ID)
}

func (r *AttendanceRepository) GetAttendanceByStudent(ctx context.Context, studentID int) ([]models.Attendance, error) {
	var attendances []models.Attendance
	query := `SELECT id, student_id, lesson_id, present, date 
	          FROM attendances WHERE student_id = $1`
	err := r.db.SelectContext(ctx, &attendances, query, studentID)
	return attendances, err
}
