package services

import (
	"context"
	"online-school/internal/models"
	"online-school/internal/repositories"
)

type AdminService struct {
	userRepo     *repositories.UserRepository
	teacherRepo  *repositories.TeacherRepository
	studentRepo  *repositories.StudentRepository
	lessonRepo   *repositories.LessonRepository
	scheduleRepo *repositories.ScheduleRepository
	paymentRepo  *repositories.PaymentRepository
}

func NewAdminService(
	userRepo *repositories.UserRepository,
	teacherRepo *repositories.TeacherRepository,
	studentRepo *repositories.StudentRepository,
	lessonRepo *repositories.LessonRepository,
	scheduleRepo *repositories.ScheduleRepository,
	paymentRepo *repositories.PaymentRepository,
) *AdminService {
	return &AdminService{
		userRepo:     userRepo,
		teacherRepo:  teacherRepo,
		studentRepo:  studentRepo,
		lessonRepo:   lessonRepo,
		scheduleRepo: scheduleRepo,
		paymentRepo:  paymentRepo,
	}
}

// Добавленные методы
func (s *AdminService) CreateTeacher(ctx context.Context, teacher *models.Teacher) error {
	return s.teacherRepo.CreateTeacher(ctx, teacher)
}

func (s *AdminService) ApproveTeacher(ctx context.Context, id int) error {
	return s.teacherRepo.ApproveTeacher(ctx, id)
}

func (s *AdminService) CreateStudent(ctx context.Context, student *models.Student) error {
	return s.studentRepo.CreateStudent(ctx, student)
}

func (s *AdminService) CreateLesson(ctx context.Context, lesson *models.Lesson) error {
	return s.lessonRepo.CreateLesson(ctx, lesson)
}

func (s *AdminService) CreateSchedule(ctx context.Context, schedule *models.Schedule) error {
	return s.scheduleRepo.CreateSchedule(ctx, schedule)
}

func (s *AdminService) CreatePayment(ctx context.Context, payment *models.Payment) error {
	return s.paymentRepo.CreatePayment(ctx, payment)
}
