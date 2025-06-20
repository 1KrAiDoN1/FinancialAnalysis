package services

import (
	"context"
	"crypto/sha1"
	"errors"
	"finance/internal/dto"
	"finance/internal/models"
	"finance/internal/repositories"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

const (
	JWTokenTTL      = 24 * time.Hour
	RefreshTokenTTL = 30 * 24 * time.Hour
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
	if req.Password != req.ConfirmPassword {
		return nil, errors.New("password and confirm password do not match")
	}
	exists, err := a.repo.UserExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check user existence: %w", err)
	}
	if exists {
		return nil, errors.New("user with this email already exists")
	}
	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	user := &models.User{
		FirstName:          req.FirstName,
		LastName:           req.LastName,
		Email:              req.Email,
		Password:           hashedPassword,
		TimeOfRegistration: time.Now(),
	}
	createdUser, err := a.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return &dto.UserInfo{
		Email:     createdUser.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}, nil

}

func (a *AuthService) SignIn(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error) {
	hashPassword, err := HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	user, err := a.repo.CheckUserVerification(ctx, req.Email, hashPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	accesstoken, err := a.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}
	refreshToken, err := a.GenerateRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return &dto.AuthResponse{
		AccessToken:  accesstoken.AccessToken,
		RefreshToken: refreshToken.RefreshToken,
		User: dto.UserInfo{
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}, nil
}

func (a *AuthService) Logout(ctx context.Context, req dto.LogoutRequest) error {
	return nil
}

func (a *AuthService) GenerateRefreshToken() (dto.RefreshTokenRequest, error) {
	refresh_token := make([]byte, 32)
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	if _, err := r.Read(refresh_token); err != nil {
		return dto.RefreshTokenRequest{}, err
	}
	return dto.RefreshTokenRequest{}, nil
}

func (a *AuthService) GenerateAccessToken(userID int) (dto.AccessTokenRequest, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   strconv.Itoa(userID),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(JWTokenTTL)),
	})
	err := godotenv.Load("/Users/pavelvasilev/Desktop/FinancialAnalysis/internal/storages/database/hash.env")
	if err != nil {
		log.Fatal(err)
		return dto.AccessTokenRequest{}, fmt.Errorf("failed to load environment file: %w", err)
	}
	secretSignInKey := os.Getenv("SECRET_SIGNINKEY")
	if secretSignInKey == "" {
		return dto.AccessTokenRequest{}, fmt.Errorf("SECRET_SIGNINKEY environment variable is not set")
	}
	tokenString, err := token.SignedString([]byte(secretSignInKey))
	if err != nil {
		return dto.AccessTokenRequest{}, fmt.Errorf("failed to sign token: %w", err)
	}
	return dto.AccessTokenRequest{
		AccessToken: tokenString,
	}, nil
}

func (a *AuthService) ValidateToken(ctx context.Context, req dto.AccessTokenRequest) (*dto.UserInfo, error) {
	return &dto.UserInfo{}, nil
}

func HashPassword(Password string) (string, error) {
	hash := sha1.New()
	err := godotenv.Load("/Users/pavelvasilev/Desktop/FinancialAnalysis/internal/storages/database/hash.env")
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	secretString := os.Getenv("SECRET_HASH") // получаем значение из файла конфигурации
	_, err1 := hash.Write([]byte(Password))
	if err1 != nil {
		log.Println("Ошибка при шифровании пароля", err)
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum([]byte(secretString))), nil

}
