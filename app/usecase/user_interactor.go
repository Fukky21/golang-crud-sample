package usecase

import (
	"example.com/m/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Users() (users *domain.Users, err error) {
	users, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) User(id int) (user *domain.User, err error) {
	user, err = interactor.UserRepository.FindById(id)
	return
}

func (interactor *UserInteractor) Create(u *domain.User) (user *domain.User, err error) {
	user, err = interactor.UserRepository.Create(u)
	return
}

func (interactor *UserInteractor) Update(id int, u *domain.User) (user *domain.User, err error) {
	user, err = interactor.UserRepository.Update(id, u)
	return
}

func (interactor *UserInteractor) Destroy(id int) (user *domain.User, err error) {
	user, err = interactor.UserRepository.Destroy(id)
	return
}