package usecase

import (
	"errors"

	"github.com/adityarizkyramadhan/hackfest-ciputra-23/api/business/repository"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/model"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/utils/uploader"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type Business struct {
	repoBusiness *repository.Business
}

func New(repoBusiness *repository.Business) *Business {
	return &Business{repoBusiness}
}

func (usecase *Business) CreateBusiness(arg *model.BusinessInput) error {
	business := new(model.Business)
	link, err := uploader.SupClient.Upload(arg.Photo)
	if err != nil {
		return err
	}
	err = copier.Copy(business, arg)
	if err != nil {
		return nil
	}
	business.ID = uuid.Must(uuid.NewV6())
	business.LinkPhoto = link
	return usecase.repoBusiness.CreateBusiness(business)
}

func (usecase *Business) CreateTestimony(arg *model.TestimonyInput, userId string) error {
	idUser, err := uuid.FromString(userId)
	if err != nil {
		return err
	}
	idBusiness, err := uuid.FromString(arg.IDBusiness)
	if err != nil {
		return err
	}
	link, err := uploader.SupClient.Upload(arg.Photo)
	if err != nil {
		return err
	}
	testimony := &model.Testimony{
		ID:         uuid.Must(uuid.NewV6()),
		IDUser:     idUser,
		IDBusiness: idBusiness,
		LinkPhoto:  link,
		Comentar:   arg.Comentar,
	}
	return usecase.repoBusiness.CreateTestimony(testimony)
}

func (usecase *Business) GetBussiness(tipe, offered string) ([]*model.Business, error) {
	businesses, err := usecase.repoBusiness.GetAllBusiness(tipe, offered)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []*model.Business{}, nil
	}
	if err != nil {
		return nil, err
	}
	return businesses, nil
}

func (usecase *Business) GetByIdBusiness(id string) (*model.Business, error) {
	return usecase.repoBusiness.GetDetailCompleteByIdBusiness(id)
}
