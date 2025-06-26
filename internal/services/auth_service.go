package services

import (
	"context"
	"errors"

	"github.com/Muhandis2000/online-school/internal/models"
	"github.com/Muhandis2000/online-school/internal/repositories"
	"github.com/Muhandis2000/online-school/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo  *repositories.UserRepository
	jwtSecret string
}

func NewAuthService(repo *repositories.Repository, jwtSecret string) *AuthService {
	return &AuthService{
		userRepo:  repo.User,
		jwtSecret: jwtSecret,
	}
}

func (s *AuthService) Register(ctx context.Context, user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	if user.Role != "student" && user.Role != "teacher" && user.Role != "admin" {
		return errors.New("invalid user role")
	}

	return s.userRepo.CreateUser(ctx, user)
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID, user.Role, s.jwtSecret)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
