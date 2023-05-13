package model

import (
	"mime/multipart"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Business struct {
	ID          uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name        string         `json:"name"`
	Region      string         `json:"region"`
	Description string         `json:"description"`
	IsAvailable bool           `json:"is_available" gorm:"default:true"`
	CloseTime   string         `json:"close_time"`
	OpenTime    string         `json:"open_time"`
	Testimonies []Testimony    `gorm:"foreignKey:IDBusiness" json:"testimonies"`
}

type BusinessInput struct {
	Name        string `json:"name" binding:"required"`
	Region      string `json:"region" binding:"required"`
	Description string `json:"description" binding:"required"`
	IsAvailable bool   `json:"is_available" binding:"required"`
	CloseTime   string `json:"close_time" binding:"required"`
	OpenTime    string `json:"open_time" binding:"required"`
}

type Testimony struct {
	ID         uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	IDBusiness uuid.UUID      `gorm:"type:char(36)" json:"id_business"`
	IDUser     uuid.UUID      `json:"id_user"`
	LinkPhoto  string         `json:"link_photo"`
	Comentar   string         `json:"comentar"`
	User       User           `gorm:"foreignKey:IDUser" json:"user"`
}

type TestimonyInput struct {
	Photo    *multipart.FileHeader `form:"link_photo" binding:"required"`
	Comentar string                `form:"comentar" binding:"required"`
}
