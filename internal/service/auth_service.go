package service

import (
	"finance/internal/models"
	"finance/internal/repository"
)

type AuthService struct {
	repos repository.Authorization
}

func NewAuthservice(repository repository.Authorization) *AuthService { // принимает какую-то структуру, которая удовлетворяет интерфейсу
	return &AuthService{
		repos: repository,
	}

}

func (r *AuthService) CreateUser(user models.User) (int, error) {
	return r.repos.CreateUser(user)
}
