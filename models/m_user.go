package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email			string		`gorm:"type:varchar(255);unique_index" validate:"required,email"`
	Password		string		`gorm:"type:varchar(255)" validate:"required" json:"-"`
	FirstName		string		`gorm:"type:varchar(255)"`
	LastName		string		`gorm:"type:varchar(255)"`
	AvatarUrl		string		`gorm:"type:varchar(255)"`
	CorverUrl		string
	Phone			string		`gorm:"type:varchar(31)"`
	Class 			string		`gorm:"type:varchar(31)"`
	School			string		`gorm:"type:varchar(63)"`
	Birthday     	time.Time
	Gender			uint
	CreatorId		uint		``
	Description		string
	ProvinceId		uint
	ExamsNumber		uint
	TestedNumber	uint
	LessonsLearnedNumber		uint
	SocialId		string
	ActivityNumber	uint
	SocialAccessToken	string		`json:"-" form:"-"`
	IsAdmin			bool		`json:"IsAdmin,omitempty"`
}

type UserFilter struct {
	Page
	Id 			uint		`form:"ID"`
	SocialId	string
}

func (f *UserFilter) GetList(db *gorm.DB) (users []User,err error) {
	offset := f.getOffset()
	err = db.Limit(f.Rows).Offset(offset).Order(f.getOrderBy()).Find(&users).Error
	if f.Count && err == nil {
		err = db.Model(&User{}).Count(&f.Total).Error
	}
	return
}

func(f UserFilter) GetA(db *gorm.DB) (user User,err error) {
	if f.Id > 0 {
		err = db.First(&user,f.Id).Error
		return
	}
	if f.SocialId != "" {
		err = db.Where("social_id = ?",f.SocialId).First(&user).Error
		return
	}
	return
}

func(f UserFilter) Delete(db *gorm.DB) error {
	return db.Where("id = ?",f.Id).Delete(&User{}).Error
}