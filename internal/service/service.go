package service

import (
	"finance/internal/models"
	"finance/internal/repository"
)

type Authorization interface { // интерфейс для методов авторизации пользователя
	CreateUser(user models.User) (int, error)
}
type Service struct { // создаем структуру, которая будет содержать интерфейсы
	Authorization
}

func NewService(repos *repository.Repository) *Service { // через конструктор создаем образ структуры
	return &Service{
		Authorization: NewAuthservice(repos.Authorization), // передаем
	}
}
