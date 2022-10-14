package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	ID        uint64    `gorm:"primarykey" json:"id"`
	Title     string    `gorm:"not null" json:"title" form:"title" valid:"required~title is required"`
	Caption   string    `json:"caption" form:"caption"`
	PhotoURL  string    `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~title is required"`
	UserID    uint64    `json:"user_id" form:"user_id"`
	User      User      `json:"user" form:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return err
	}

	err = nil
	return
}