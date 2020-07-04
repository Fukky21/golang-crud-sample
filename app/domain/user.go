package domain

import (
	"github.com/jinzhu/gorm"
	"errors"
)

type User struct {
  gorm.Model
  Name string
	Age int
	Sex int // 1: 男性, 2: 女性, 3: その他
	Email string
}

type Users []User

func (user *User) UserValidation() error {
	// Name
	if len(user.Name) <= 0 {
		return errors.New("Name Invalid!")
	}
	// Age
	if user.Age <= 0 {
		return errors.New("Age Invalid!")
	}
	// Sex
	if user.Sex <= 0 || user.Sex > 3 {
		return errors.New("Sex Invalid!")
	}
	// Email
	if len(user.Email) <= 0 {
		return errors.New("Email Invalid!")
	}
	return nil
}