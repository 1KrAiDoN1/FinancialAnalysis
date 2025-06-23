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
	userprofile, err := s.repo.GetProfile(ctx, userID)
	if err != nil {
		return dto.UserProfile{}, err
	}
	res_profile := dto.UserProfile{
		Email:     userprofile.Email,
		FirstName: userprofile.FirstName,
		LastName:  userprofile.LastName,
		CreatedAt: userprofile.TimeOfRegistration,
	}
	return res_profile, nil
}

func (s *UserService) DeleteAccount(ctx context.Context, userID uint) error {
	return s.repo.DeleteUser(ctx, userID)
}

func (s *UserService) GetUserStats(ctx context.Context, userID uint) (dto.UserStats, error) {
	userstats, err := s.repo.GetUserStats(ctx, userID)
	if err != nil {
		return dto.UserStats{}, err
	}
	res_stats := dto.UserStats{
		TotalExpenses:   userstats.TotalExpenses,
		TotalCategories: userstats.TotalCategories,
		TotalBudgets:    userstats.TotalBudgets,
		MonthlyExpenses: userstats.MonthlyExpenses,
		WeeklyExpenses:  userstats.WeeklyExpenses,
		TopCategories:   nil,
	}
	return res_stats, nil
}
