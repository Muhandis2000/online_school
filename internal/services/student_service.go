package services

import (
	"context"
	"online-school/internal/models"
	"online-school/internal/repositories"
)

type StudentService struct {
	studentRepo    *repositories.StudentRepository
	scheduleRepo   *repositories.ScheduleRepository
	lessonRepo     *repositories.LessonRepository
	homeworkRepo   *repositories.HomeworkRepository
	gradeRepo      *repositories.GradeRepository
	attendanceRepo *repositories.AttendanceRepository
}

func NewStudentService(
	studentRepo *repositories.StudentRepository,
	scheduleRepo *repositories.ScheduleRepository,
	lessonRepo *repositories.LessonRepository,
	homeworkRepo *repositories.HomeworkRepository,
	gradeRepo *repositories.GradeRepository,
	attendanceRepo *repositories.AttendanceRepository,
) *StudentService {
	return &StudentService{
		studentRepo:    studentRepo,
		scheduleRepo:   scheduleRepo,
		lessonRepo:     lessonRepo,
		homeworkRepo:   homeworkRepo,
		gradeRepo:      gradeRepo,
		attendanceRepo: attendanceRepo,
	}
}

func (s *StudentService) GetSchedule(ctx context.Context, studentID int) ([]models.Schedule, error) {
	student, err := s.studentRepo.GetStudentByID(ctx, studentID)
	if err != nil {
		return nil, err
	}
	return s.scheduleRepo.GetScheduleForGroup(ctx, student.Group)
}

func (s *StudentService) GetLessons(ctx context.Context, studentID int) ([]models.Lesson, error) {
	student, err := s.studentRepo.GetStudentByID(ctx, studentID)
	if err != nil {
		return nil, err
	}
	return s.lessonRepo.GetLessonsForGroup(ctx, student.Group)
}

func (s *StudentService) SubmitHomework(ctx context.Context, homework *models.Homework) error {
	return s.homeworkRepo.CreateHomework(ctx, homework)
}

func (s *StudentService) GetGrades(ctx context.Context, studentID int) ([]models.Grade, error) {
	return s.gradeRepo.GetGradesByStudent(ctx, studentID)
}

func (s *StudentService) GetAttendance(ctx context.Context, studentID int) ([]models.Attendance, error) {
	return s.attendanceRepo.GetAttendanceByStudent(ctx, studentID)
}
