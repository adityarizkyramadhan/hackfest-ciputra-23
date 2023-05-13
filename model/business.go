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
	Type        string         `json:"type"`
	Offered     string         `json:"offered"`
	Latitude    float64        `json:"latitude" binding:"required"`
	Longitude   float64        `json:"longitude" binding:"required"`
	LinkPhoto   string         `json:"link_photo"`
	Testimonies []Testimony    `gorm:"foreignKey:IDBusiness" json:"testimonies"`
}

type BusinessInput struct {
	Name        string                `form:"name" binding:"required"`
	Region      string                `form:"region" binding:"required"`
	Description string                `form:"description" binding:"required"`
	CloseTime   string                `form:"close_time" binding:"required"`
	OpenTime    string                `form:"open_time" binding:"required"`
	Type        string                `form:"type" binding:"required"`
	Offered     string                `form:"offered" binding:"required"`
	Latitude    float64               `form:"latitude" binding:"required"`
	Longitude   float64               `form:"longitude" binding:"required"`
	Photo       *multipart.FileHeader `form:"photo" binding:"required"`
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
	Photo      *multipart.FileHeader `form:"photo" binding:"required"`
	IDBusiness string                `form:"id_business" binding:"required"`
	Comentar   string                `form:"comentar" binding:"required"`
}
