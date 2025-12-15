package service

import (
	"context"
	"time"

	"github.com/student/go-user-api/db/sqlc"
	"github.com/student/go-user-api/internal/models"
	"github.com/student/go-user-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CalculateAge calculates age from date of birth
// Corrects age if birthday hasn't occurred yet this year
func CalculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()

	// Check if birthday hasn't occurred yet this year
	if now.Month() < dob.Month() || (now.Month() == dob.Month() && now.Day() < dob.Day()) {
		age--
	}

	return age
}

func (s *UserService) CreateUser(ctx context.Context, req models.CreateUserRequest) (*models.UserResponse, error) {
	dob, err := models.ParseDOB(req.DOB)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.Create(ctx, req.Name, dob)
	if err != nil {
		return nil, err
	}

	return s.toUserResponse(user), nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int32) (*models.UserResponse, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.toUserResponse(user), nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (*models.UserResponse, error) {
	dob, err := models.ParseDOB(req.DOB)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.Update(ctx, id, req.Name, dob)
	if err != nil {
		return nil, err
	}

	return s.toUserResponse(user), nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}

func (s *UserService) ListUsers(ctx context.Context) ([]models.UserResponse, error) {
	users, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]models.UserResponse, len(users))
	for i, user := range users {
		responses[i] = *s.toUserResponse(user)
	}

	return responses, nil
}

func (s *UserService) toUserResponse(user sqlc.User) *models.UserResponse {
	return &models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  models.FormatDOB(user.Dob),
		Age:  CalculateAge(user.Dob),
	}
}
