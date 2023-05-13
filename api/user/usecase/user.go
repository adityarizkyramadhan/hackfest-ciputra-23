package usecase

import (
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/api/user/repository"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/middleware"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/model"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/copier"
)

type User struct {
	repoUser *repository.User
}

func New(repoUser *repository.User) *User {
	return &User{repoUser}
}

func (usecase *User) Register(arg *model.UserRegister) (string, error) {
	user := new(model.User)
	err := copier.Copy(user, arg)
	if err != nil {
		return "", err
	}
	user.ID = uuid.Must(uuid.NewV6())
	if err := usecase.repoUser.Create(user); err != nil {
		return "", err
	}
	token, err := middleware.GenerateJWToken(user.ID.String())
	if err != nil {
		return "", err
	}
	return token, err
}

func (usecase *User) Login(arg *model.UserLogin) (string, error) {
	user, err := usecase.repoUser.FindByNumberPhone(arg.PhoneNumber)
	if err != nil {
		return "", err
	}
	token, err := middleware.GenerateJWToken(user.ID.String())
	if err != nil {
		return "", err
	}
	return token, err
}
