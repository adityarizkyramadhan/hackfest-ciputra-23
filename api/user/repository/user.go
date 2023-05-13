package repository

import (
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/model"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func New(db *gorm.DB) *User {
	return &User{db}
}

func (rs *User) Create(arg *model.User) error {
	return rs.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(arg).Error
	})
}

func (rs *User) FindById(id string) (*model.User, error) {
	var user model.User
	if err := rs.db.Where("id = ?", id).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (rs *User) FindByNumberPhone(phoneNumber string) (*model.User, error) {
	var user model.User
	if err := rs.db.Where("phone_number = ?", phoneNumber).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
