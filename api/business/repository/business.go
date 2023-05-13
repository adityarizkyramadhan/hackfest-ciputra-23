package repository

import (
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/model"
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
	query := repo.db
	if tipe != "" {
		query = query.Where("type = ?", tipe)
	}
	if offered != "" {
		query = query.Where("offered = ?", offered)
	}

	if err := query.Find(&businesses).Error; err != nil {
		return nil, err
	}
	return businesses, nil
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
