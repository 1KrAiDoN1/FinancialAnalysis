package services

import (
	"context"
	"finance/internal/dto"
	"finance/internal/repositories"
)

type AuthService struct {
	repo repositories.AuthRepositoryInterface
}

func NewAuthService(repo repositories.AuthRepositoryInterface) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (a *AuthService) SignUp(ctx context.Context, req dto.RegisterRequest) (*dto.UserInfo, error) {
	return &dto.UserInfo{}, nil
}

func (a *AuthService) SignIn(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error) {
	return &dto.AuthResponse{}, nil
}

func (a *AuthService) Logout(ctx context.Context, req dto.LogoutRequest) error {
	return nil
}

func (a *AuthService) GenerateRefreshToken() (*dto.AuthResponse, error) {
	return &dto.AuthResponse{}, nil
}

func (a *AuthService) GenerateAccessToken(req dto.LoginRequest) (*dto.AuthResponse, error) {
	return &dto.AuthResponse{}, nil
}

func (a *AuthService) ValidateToken(ctx context.Context, req dto.AccessTokenRequest) (*dto.UserInfo, error) {
	return &dto.UserInfo{}, nil
}
