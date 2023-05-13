package repository

import (
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/model"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/utils/customserror"
	"gorm.io/gorm"
)

type Business struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Business {
	return &Business{db}
}

func (repo *Business) CreateBusiness(arg *model.Business) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(arg).Error
	})
}

func (repo *Business) CreateTestimony(arg *model.Testimony) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(arg).Error
	})
}

func (repo *Business) GetAllBusiness(tipe, offered string) ([]*model.Business, error) {
	var businesses []*model.Business
	if tipe != "" && offered != "" {
		if err := repo.db.Where("offered = ?", offered).Where("type = ?", tipe).Find(&businesses).Error; err != nil {
			return nil, err
		}
		return businesses, nil
	}
	if offered != "" {
		if err := repo.db.Where("offered = ?", offered).Find(&businesses).Error; err != nil {
			return nil, err
		}
		return businesses, nil
	}
	if tipe != "" {
		if err := repo.db.Where("type = ?", tipe).Find(&businesses).Error; err != nil {
			return nil, err
		}
		return businesses, nil
	}
	return nil, customserror.ErrNoQuery
}

func (repo *Business) GetByIdBusiness(id string) (*model.Business, error) {
	var business *model.Business
	if err := repo.db.Where("id = ?", id).Take(&business).Error; err != nil {
		return nil, err
	}
	return business, nil
}

func (repo *Business) GetAllTestimony(idBusiness string) ([]*model.Testimony, error) {
	var testimonies []*model.Testimony
	if err := repo.db.Where("id_business = ?", idBusiness).Find(&testimonies).Error; err != nil {
		return nil, err
	}
	return testimonies, nil
}

func (repo *Business) GetDetailCompleteByIdBusiness(id string) (*model.Business, error) {
	var business *model.Business
	if err := repo.db.Preload("Testimonies").Preload("Testimonies.User").Where("id = ?", id).Take(&business).Error; err != nil {
		return nil, err
	}
	return business, nil
}
