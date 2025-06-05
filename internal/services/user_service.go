package services

import (
	"context"
	"finance/internal/dto"
	"finance/internal/repositories"
)

type UserService struct {
	repo repositories.UserRepositoryInterface
}

func NewUserService(repo repositories.UserRepositoryInterface) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetProfile(ctx context.Context, userID uint) (*dto.UserProfile, error) {

}

func (s *UserService) UpdateProfile(ctx context.Context, userID uint, req dto.UpdateProfileRequest) error {

}

func (s *UserService) DeleteAccount(ctx context.Context, userID uint) error {

}

func (s *UserService) GetUserStats(ctx context.Context, userID uint) (*dto.UserStats, error) {

}
