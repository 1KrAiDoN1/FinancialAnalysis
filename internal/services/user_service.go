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

func (s *UserService) GetProfile(ctx context.Context, userID uint) (dto.UserProfile, error) {
	return dto.UserProfile{}, nil
}

func (s *UserService) DeleteAccount(ctx context.Context, userID uint) error {
	return nil
}

func (s *UserService) GetUserStats(ctx context.Context, userID uint) (dto.UserStats, error) {
	return dto.UserStats{}, nil
}
