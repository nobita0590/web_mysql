package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type (
	Fags struct {
		gorm.Model
		Title				string
		PrettyUrl			string
		Content				string
		CommentsNumber		int
		SubjectId			uint
		ClassId				uint
		UserId				uint
		IsDone				bool
		SubjectName			string			`gorm:"-"`
		ClassName			string			`gorm:"-"`
		UserName			string			`gorm:"-"`
		AvatarUrl			string			`gorm:"-"`
	}
	FagsFilter struct {
		ID		uint
		Page
		SubjectId 			uint
		ClassId 			uint
		Option 				string
		IsHot				bool
	}
)

func (f *FagsFilter) GetList(db *gorm.DB,userId uint) (fags []Fags,err error) {
	offset := f.getOffset()
	selectField := ""
	order := "fags.created_at desc"
	if f.SubjectId > 0 {
		db = db.Where("fags.subject_id = ?",f.SubjectId)
	}
	if f.ClassId > 0 {
		db = db.Where("fags.class_id = ?",f.ClassId)
	}
	if f.Option == "no_done" {
		//db = db.Where("fags.is_done = 0")
		db = db.Where("fags.comments_number = 0")
	}
	if f.IsHot {
		t := time.Now().AddDate(0,0,-7)
		selectField = `(select count(id) from fag_comments where fag_comments.fag_id = fags.id and fag_comments.created_at > '`+
			t.Format(`2006-01-02 15:04:05`) + `') hot_num,`
		order = "hot_num desc, " + order
	}
	if userId > 0 {
		switch f.Option {
		case "my_fag":
			db = db.Where("fags.user_id = ?",userId)
		case "my_answer":
			db = db.Joins("left join fag_comments fc ON fc.fag_id = fags.id").
				Group("fags.id").Where("fc.user_id = ?",userId)
		}
	}
	err = db.Table("fags").
		Select(selectField +
			"fags.*,class.value class_name,subject.value subject_name,CONCAT(users.last_name,' ',users.first_name) user_name,users.avatar_url").
		Joins("left join select_sources class on fags.class_id = class.id").
		Joins("left join select_sources subject on fags.subject_id = subject.id").
		Joins("left join users on users.id = fags.user_id").
		Where("fags.deleted_at IS NULL").
		Limit(f.Rows).Offset(offset).Order(order).Find(&fags).Error
	if f.Count && err == nil {
		err = db.Model(&Fags{}).Count(&f.Total).Error
	}
	return
}

func(f FagsFilter) GetA(db *gorm.DB) (fag Fags,err error) {
	if f.ID > 0 {
		db = db.Where("fags.id = ?",f.ID)
	}
	err = db.Table("fags").
		Select("fags.*,class.value class_name,subject.value subject_name,CONCAT(users.last_name,' ',users.first_name) user_name,users.avatar_url").
		Joins("left join select_sources class on fags.class_id = class.id").
		Joins("left join select_sources subject on fags.subject_id = subject.id").
		Joins("left join users on users.id = fags.user_id").
		Where("fags.deleted_at IS NULL").
		Limit(1).First(&fag).Error
	return
}

func(f FagsFilter) Delete(db *gorm.DB) error {
	return db.Where("id = ?",f.ID).Delete(&Fags{}).Error
}