package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Title     string    `gorm:"not null" json:"title" form:"title" valid:"required~title is required"`
	Caption   string    `json:"caption" form:"caption"`
	PhotoURL  string    `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~photo url is required"`
	UserID    uint      `json:"user_id" form:"user_id"`
	User      User      `json:"user" form:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Comment   []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"comment"`
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

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return err
	}

	err = nil
	return
}
