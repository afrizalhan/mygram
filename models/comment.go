package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UserID    uint      `json:"user_id" form:"user_id"`
	User      User      `json:"user" form:"user"`
	PhotoID   uint      `json:"photo_id" form:"photo_id"`
	Photo     Photo     `json:"photo" form:"photo"`
	Message   string    `gorm:"not null" json:"message" form:"message" valid:"required~message is required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return err
	}

	err = nil
	return
}
