package models

import (
	"github.com/jinzhu/gorm"
)

type(
	Document struct {
		gorm.Model
		Name 		string
		PrettyUrl	string
		PathStore	string
		Description	string
		CreatorId	uint
		CreatorName	string		`gorm:"-"`
		ClassId		uint
		ClassName	string		`gorm:"-"`
		SubjectId	uint
		SubjectName	string		`gorm:"-"`
		DownloadNumber		uint
	}
	DocumentFilter struct {
		Page
		Id  		uint	`form:"ID"`
		IsFill		bool
		ClassId		uint
		SubjectId 	uint
	}
)

func (f *DocumentFilter) GetList(db *gorm.DB) (documents []Document,err error) {
	offset := f.getOffset()
	if f.SubjectId > 0 {
		db = db.Where("documents.subject_id = ?",f.SubjectId)
	}
	if f.ClassId > 0 {
		db = db.Where("documents.class_id = ?",f.ClassId)
	}
	err = db.Table("documents").
		Select("documents.*,subject.value AS subject_name,class.value AS class_name,concat(users.last_name,' ',users.first_name) AS creator_name").
		Joins("LEFT JOIN users ON documents.creator_id = users.id").
		Joins("LEFT JOIN select_sources AS subject ON documents.subject_id = subject.id").
		Joins("LEFT JOIN select_sources AS class ON documents.class_id = class.id").
		Order("updated_at desc").
		Limit(f.Rows).Offset(offset).Find(&documents).Error

	if f.Count && err == nil {
		err = db.Model(&Document{}).Count(&f.Total).Error
	}
	return
}

func(f DocumentFilter) GetA(db *gorm.DB) (document Document,err error) {
	err = db.First(&document,f.Id).Error
	return
}

func(f DocumentFilter) GetAForFront(db *gorm.DB) (doc Document,err error) {
	err = db.Select("documents.*,subject.value AS subject_name,class.value AS class_name,concat(users.last_name,' ',users.first_name) AS creator_name").
		Joins("LEFT JOIN users ON documents.creator_id = users.id").
		Joins("LEFT JOIN select_sources AS subject ON documents.subject_id = subject.id").
		Joins("LEFT JOIN select_sources AS class ON documents.class_id = class.id").
		Where("documents.id = ?",f.Id).First(&doc).Error
	return
}

func(f DocumentFilter) Delete(db *gorm.DB) error {
	return db.Where("id = ?",f.Id).Delete(&Document{}).Error
}