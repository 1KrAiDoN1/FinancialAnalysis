package services

import (
	"context"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"finance/internal/dto"
	"finance/internal/models"
	"finance/internal/repositories"
	"fmt"
	"log"
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
		ID:        uint(createdUser.ID),
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
	user, err := a.repo.CheckUserVerification(ctx, req.Email, hashPassword) //нужно, чтобы возвращал всю информацию о пользователе(id, email, first name, last name)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
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
			ID:        uint(user.ID),
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}, nil
}

func (a *AuthService) GenerateRefreshToken() (dto.RefreshTokenRequest, error) {
	refresh_token := make([]byte, 32)
	if _, err := rand.Read(refresh_token); err != nil {
		return dto.RefreshTokenRequest{}, err
	}
	token := base64.URLEncoding.EncodeToString(refresh_token)

	return dto.RefreshTokenRequest{
		RefreshToken: token,
		ExpiresAt:    time.Now().Add(RefreshTokenTTL),
	}, nil
}

func (a *AuthService) GenerateAccessToken(userID int) (dto.AccessTokenRequest, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   strconv.Itoa(userID),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(JWTokenTTL)),
	})
	err := godotenv.Load("./internal/storages/database/hash.env")
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
		ExpiresAt:   time.Now().Add(JWTokenTTL),
	}, nil
}

func (a *AuthService) ValidateToken(ctx context.Context, req dto.AccessTokenRequest) (*dto.UserID, error) {

	err := godotenv.Load("./internal/storages/database/hash.env")
	if err != nil {
		return &dto.UserID{}, fmt.Errorf("failed to load environment file: %w", err)
	}

	secretSignInKey := os.Getenv("SECRET_SIGNINKEY")
	if secretSignInKey == "" {
		return &dto.UserID{}, fmt.Errorf("SECRET_SIGNINKEY environment variable is not set")
	}

	token, err := jwt.ParseWithClaims(req.AccessToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretSignInKey), nil
	})

	if err != nil {
		return &dto.UserID{}, fmt.Errorf("invalid token: %w", err)
	}

	// Проверяем валидность claims
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
			return &dto.UserID{}, err
		}

		userID, err := strconv.Atoi(claims.Subject)
		if err != nil {
			return &dto.UserID{}, fmt.Errorf("invalid user ID in access token: %w", err)
		}

		return &dto.UserID{
			UserID: userID,
		}, nil
	}

	return nil, fmt.Errorf("invalid token claims")
}

func HashPassword(Password string) (string, error) {
	hash := sha1.New()
	err := godotenv.Load("./internal/storages/database/hash.env")
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

func (a *AuthService) GetUserIDbyRefreshToken(ctx context.Context, refresh_token string) (int, error) {
	return a.repo.GetUserIDbyRefreshToken(ctx, refresh_token)
}

func (a *AuthService) RemoveOldRefreshToken(ctx context.Context, userID int) error {
	return a.repo.RemoveOldRefreshToken(ctx, userID)
}

func (a *AuthService) SaveNewRefreshToken(ctx context.Context, user_id int, token dto.RefreshTokenRequest) error {
	refresh_token := models.RefreshToken{
		Token:     token.RefreshToken,
		ExpiresAt: token.ExpiresAt,
	}
	return a.repo.SaveNewRefreshToken(ctx, user_id, refresh_token)
}
