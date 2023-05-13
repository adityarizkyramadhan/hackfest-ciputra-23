package usecase

import (
	"errors"

	"github.com/adityarizkyramadhan/hackfest-ciputra-23/api/user/repository"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/middleware"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/model"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
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

func (usecase *User) AddLocation(arg *model.UserLocationRequest, userId string) error {
	location := new(model.UserLocation)
	userIdUUID, err := uuid.FromString(userId)
	if err != nil {
		return err
	}
	err = copier.Copy(location, arg)
	if err != nil {
		return err
	}
	location.ID = uuid.Must(uuid.NewV6())
	location.UserID = userIdUUID
	return usecase.repoUser.AddLocation(location)
}

func (usecase *User) IsUserAddLocation(userId string) (bool, error) {
	_, err := usecase.repoUser.FindLocationUser(userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func (usecase *User) GetUserLocation(userId string) (*model.UserLocation, error) {
	return usecase.repoUser.FindLocationUser(userId)
}
