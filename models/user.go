package models

import (
	"errors"
	"mygram/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID        uint          `gorm:"primarykey" json:"id"`
	Username  string        `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~your username is required"`
	Email     string        `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~email is required,email~Invalid email format"`
	Password  string        `gorm:"not null" json:"password" form:"password" valid:"required~password is required,minstringlength(6)~password must be at least 6 characters"`
	Age       int           `gorm:"not null" json:"age" form:"age" valid:"required~age is required"`
	CreatedAt time.Time     `json:"created_at,omitempty"`
	UpdatedAt time.Time     `json:"updated_at,omitempty"`
	Photo     []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photo"`
	Comment   []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comment"`
	Socmed    []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"socmed"`
}

type UserUpdate struct {
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~email is required,email~Invalid email format"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~password is required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return err
	}

	if u.Age < 8 {
		err = errors.New("you must be at least 8 years old")
		return err
	}

	u.Password = helpers.HashPassword(u.Password)

	err = nil
	return
}

func (up *UserUpdate) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(up)

	if errCreate != nil {
		err = errCreate
		return err
	}

	err = nil
	return
}
