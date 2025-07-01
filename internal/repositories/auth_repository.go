package repositories

import (
	"context"
	"finance/internal/models"
	storage "finance/internal/storages"
)

type AuthRepository struct {
	storage storage.AuthStorageInterface
}

func NewAuthRepository(storage storage.AuthStorageInterface) *AuthRepository { //конструктор
	return &AuthRepository{
		storage: storage,
	}
}

func (r *AuthRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	query := `INSERT INTO users (first_name, last_name, email, password, time_of_registration) VALUES ($1, $2, $3, $4, $5) RETURNING id, email, first_name, last_name`
	result, err := r.storage.CreateUser(ctx, query, user.FirstName, user.LastName, user.Email, user.Password, user.TimeOfRegistration)
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:        result.ID,
		Email:     result.Email,
		FirstName: result.FirstName,
		LastName:  result.LastName,
	}, nil
}

func (r *AuthRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return &models.User{}, nil // логика для получения пользователя по email
}

func (r *AuthRepository) CheckUserVerification(ctx context.Context, email string, hashpassword string) (*models.User, error) {
	return &models.User{}, nil // логика для проверки пользователя
}
func (r *AuthRepository) GetUserByID(ctx context.Context, userID uint) (*models.User, error) {
	return &models.User{}, nil // логика для получения пользователя по id
}

func (r *AuthRepository) UserExistsByEmail(ctx context.Context, email string) (bool, error) {
	return false, nil
}

func (r *AuthRepository) GetUserIDbyRefreshToken(refresh_token string) (int, error) {
	return 0, nil
}
func (r *AuthRepository) RemoveOldRefreshToken(userID int) error {
	return nil
}
func (r *AuthRepository) SaveNewRefreshToken(token models.RefreshToken) error {
	return nil
}
