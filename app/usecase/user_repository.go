package usecase

import (
	"example.com/m/domain"
)

type UserRepository interface {
	FindAll() (*domain.Users, error)
	FindById(int) (*domain.User, error)
	Create(*domain.User) (*domain.User, error)
	Update(int, *domain.User) (*domain.User, error)
	Destroy(int) (*domain.User, error)
}