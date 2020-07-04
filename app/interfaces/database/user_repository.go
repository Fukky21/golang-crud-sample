package database

import (
	"example.com/m/domain"
	"errors"
)

type UserRepository struct {
	SqlHandler SqlHandler
}

func (repo *UserRepository) FindAll() (*domain.Users, error) {
	result, err := repo.SqlHandler.Find(&domain.Users{})

	if err != nil {
		return nil, err
	}

	if users, ok := result.(*domain.Users); ok {
		return users, nil
	} else {
		return nil, errors.New("FindAll Successed, But Return Users Failed!")
	}
}

func (repo *UserRepository) FindById(id int) (*domain.User, error) {
	result, err := repo.SqlHandler.FindById(&domain.User{}, id)

	if err != nil {
		return (*domain.User)(nil), err
	}

	if user, ok := result.(*domain.User); ok {
		return user, nil
	} else {
		return (*domain.User)(nil), errors.New("FindById Successed, But Return User Failed!")
	}
}

func (repo *UserRepository) Create(u *domain.User) (*domain.User, error) {
	result, err := repo.SqlHandler.Create(u)
	
	if err != nil {
		return (*domain.User)(nil), err
	}

	if user, ok := result.(*domain.User); ok {
		return user, nil
	} else {
		return (*domain.User)(nil), errors.New("Create Successed, But Return User Failed!")
	}
}

func (repo *UserRepository) Update(id int, u *domain.User) (*domain.User, error) {
	findByIdResult, findByIdErr := repo.SqlHandler.FindById(&domain.User{}, id)

	if findByIdErr != nil {
		return (*domain.User)(nil), findByIdErr
	}

	targetUser := findByIdResult.(*domain.User)
	targetUser.Name = u.Name
	targetUser.Age = u.Age
	targetUser.Sex = u.Sex
	targetUser.Email = u.Email

	saveResult, saveErr := repo.SqlHandler.Save(targetUser)

	if saveErr != nil {
		return (*domain.User)(nil), saveErr
	}

	if user, ok := saveResult.(*domain.User); ok {
		return user, nil
	} else {
		return (*domain.User)(nil), errors.New("Update Successed, But Return User Failed!")
	}
}

func (repo *UserRepository) Destroy(id int) (*domain.User, error) {
	findByIdResult, findByIdErr := repo.SqlHandler.FindById(&domain.User{}, id)

	if findByIdErr != nil {
		return (*domain.User)(nil), findByIdErr
	}

	deleteResult, deleteErr := repo.SqlHandler.Delete(findByIdResult)

	if deleteErr != nil {
		return (*domain.User)(nil), deleteErr
	}

	if user, ok := deleteResult.(*domain.User); ok {
		return user, nil
	} else {
		return (*domain.User)(nil), errors.New("Delete Successed, But Return User Failed!")
	}
}