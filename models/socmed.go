package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             uint    `gorm:"primarykey" json:"id"`
	Name           string    `gorm:"not null" json:"name" form:"name" valid:"required~name is required"`
	SocialMediaUrl string    `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~social media url is required"`
	UserID         uint64    `json:"user_id" form:"user_id"`
	User           User      `json:"user" form:"user"`
}

func (sm *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(sm)

	if errCreate != nil {
		err = errCreate
		return err
	}

	err = nil
	return
}
