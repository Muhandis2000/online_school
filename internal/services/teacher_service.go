package services

import (
	"context"
	"online-school/internal/models"
	"online-school/internal/repositories"
)

type TeacherService struct {
	teacherRepo    *repositories.TeacherRepository
	lessonRepo     *repositories.LessonRepository
	scheduleRepo   *repositories.ScheduleRepository
	attendanceRepo *repositories.AttendanceRepository
	gradeRepo      *repositories.GradeRepository
}

func NewTeacherService(
	teacherRepo *repositories.TeacherRepository,
	lessonRepo *repositories.LessonRepository,
	scheduleRepo *repositories.ScheduleRepository,
	attendanceRepo *repositories.AttendanceRepository,
	gradeRepo *repositories.GradeRepository,
) *TeacherService {
	return &TeacherService{
		teacherRepo:    teacherRepo,
		lessonRepo:     lessonRepo,
		scheduleRepo:   scheduleRepo,
		attendanceRepo: attendanceRepo,
		gradeRepo:      gradeRepo,
	}
}

func (s *TeacherService) CreateLesson(ctx context.Context, lesson *models.Lesson) error {
	return s.lessonRepo.CreateLesson(ctx, lesson)
}

func (s *TeacherService) RecordAttendance(ctx context.Context, attendance *models.Attendance) error {
	return s.attendanceRepo.CreateAttendance(ctx, attendance)
}

func (s *TeacherService) RecordGrade(ctx context.Context, grade *models.Grade) error {
	return s.gradeRepo.CreateGrade(ctx, grade)
}

func (s *TeacherService) GetLessonsByTeacher(ctx context.Context, teacherID int) ([]models.Lesson, error) {
	return s.lessonRepo.GetLessonsByTeacher(ctx, teacherID)
}

func (s *TeacherService) GetScheduleByTeacher(ctx context.Context, teacherID int) ([]models.Schedule, error) {
	return s.scheduleRepo.GetScheduleByTeacher(ctx, teacherID)
}
