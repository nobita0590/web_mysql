package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type (
	CourseRegister struct {
		ID 			uint
		CourseId	uint
		UserId		uint
		CreatedAt	time.Time
		CourseName	string		`gorm:"-"`
		UserName	string		`gorm:"-"`
		UserPhone	string		`gorm:"-"`
		UserEmail	string		`gorm:"-"`
	}
	CourseRegisterFilter struct {
		Page
		ID 			uint
		CourseId	uint
		UserId		uint
	}
)

func (f *CourseRegisterFilter) GetARegis(db *gorm.DB) (CourseRegister,error){
	courseRegister := CourseRegister{}
	err := db.Where("course_id = ?", f.CourseId).Where("user_id = ?",f.UserId).First(&courseRegister).Error
	return courseRegister,err
}

func (f *CourseRegisterFilter) GetList(db *gorm.DB) (courseRegisters []CourseRegister,err error) {
	if f.CourseId > 0 {
		db = db.Where("course_registers.course_id = ?",f.CourseId)
	}
	if f.UserId > 0 {
		db = db.Where("course_registers.user_id = ?",f.UserId)
	}
	offset := f.getOffset()
	err = db.Table("course_registers").
		Select(`course_registers.*,courses.title course_name,users.phone user_phone,users.email user_email,
			CONCAT(users.last_name,' ',users.first_name) user_name`).
		Joins("left join courses on course_registers.course_id = courses.id").
		Joins("left join users on users.id = course_registers.user_id").
		Limit(f.Rows).Offset(offset).Order("course_registers.created_at desc").Find(&courseRegisters).Error
	if f.Count && err == nil {
		err = db.Model(&CourseRegister{}).Count(&f.Total).Error
	}
	return
}