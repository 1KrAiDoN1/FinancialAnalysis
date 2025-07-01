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

func (r *AuthRepository) CheckUserVerification(ctx context.Context, email string, hashpassword string) (models.User, error) {
	query := `SELECT id, first_name, last_name, email FROM users WHERE email = $1 AND password = $2`
	result, err := r.storage.CheckUserVerification(ctx, query, email, hashpassword)
	if err != nil {
		return models.User{}, err
	}
	return models.User{
		ID:        result.ID,
		Email:     result.Email,
		FirstName: result.FirstName,
		LastName:  result.LastName,
	}, nil
}

func (r *AuthRepository) UserExistsByEmail(ctx context.Context, email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	result, err := r.storage.UserExistsByEmail(ctx, query, email)
	if err != nil {
		return true, err
	}
	return result, nil

}

func (r *AuthRepository) GetUserIDbyRefreshToken(ctx context.Context, refresh_token string) (int, error) {
	query := `SELECT user_id FROM refresh_tokens WHERE token = $1 AND expires_at > NOW()`
	user_id, err := r.storage.GetUserIDbyRefreshToken(ctx, query, refresh_token)
	if err != nil {
		return 0, err
	}
	return user_id, nil
}
func (r *AuthRepository) RemoveOldRefreshToken(ctx context.Context, userID int) error {
	query := `DELETE FROM refresh_tokens WHERE user_id = $1`
	err := r.storage.RemoveOldRefreshToken(ctx, query, userID)
	if err != nil {
		return err
	}
	return nil
}
func (r *AuthRepository) SaveNewRefreshToken(ctx context.Context, user_id int, token models.RefreshToken) error {
	query := `INSERT INTO refresh_tokens (user_id, token, expires_at) VALUES ($1, $2, $3)`
	err := r.storage.SaveNewRefreshToken(ctx, query, user_id, token)
	if err != nil {
		return err
	}
	return nil
}
