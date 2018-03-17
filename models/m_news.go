package models

import (
	"github.com/jinzhu/gorm"
)

type (
	News struct {
		gorm.Model
		Title		string		`gorm:"type:varchar(255)" validate:"required"`
		PrettyUrl	string		`gorm:"type:varchar(255);unique_index" validate:"required"`
		Description	string		`gorm:"type:varchar(1023)"`
		Content		string		`validate:"required"`
		ImageUrl	string		``
		CreatorId	uint
		Creator		User		`json:",omitempty" validate:"-"`
		CategoryId	uint
		Views		uint		`form:"-"`
	}
	NewsCategory struct {
		ID 		uint
		Title		string		`gorm:"type:varchar(255)" validate:"required"`
		PrettyUrl	string		`gorm:"type:varchar(255)" validate:"required"`
		Description	string
	}
	NewsFilter struct {
		Page
		Id		uint		`form:"ID"`
		From	uint
		To 		uint
		Title	string
		CategoriesId	[]uint
		IgnoreIds		[]uint
	}
	NewsCategoryFilter struct {
		Page
		Id		uint		`form:"ID"`
	}
)

func (f *NewsFilter) GetList(db *gorm.DB) (listNews []News,err error) {
	offset := f.getOffset()
	// err = db.Model(&News{}).Limit(f.Rows).Offset(offset).Find(&listNews)/*.Related(&User{})*/.Error
	if len(f.Title) > 0 {
		db = db.Where("news.title LIKE ?","%"+f.Title+"%")
	}
	if len(f.CategoriesId) > 0 {
		db = db.Where("news.category_id IN (?)",f.CategoriesId)
	}
	if len(f.IgnoreIds) > 0 {
		db = db.Where("news.id NOT IN (?)",f.IgnoreIds)
	}
	rows,err := db.Table("news").Select(`news.id,news.title,news.pretty_url,news.creator_id,news.image_url,news.created_at,news.updated_at,news.description,news.category_id,
		users.email,users.first_name,users.last_name`).
		Joins("LEFT JOIN users ON news.creator_id = users.id").
		Order("news.created_at desc").
		Where("news.deleted_at IS NULL").
		Order("news.updated_at desc").
		Limit(f.Rows).Offset(offset).Rows()
	news := News{}
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		news = News{}
		err = db.ScanRows(rows, &news)
		if err != nil {
			return
		}
		err = db.ScanRows(rows, &(news.Creator))
		if err != nil {
			return
		}
		news.Creator.ID = news.CreatorId
		news.Creator.CreatorId = 0
		listNews = append(listNews,news)

	}
	if f.Count && err == nil {
		err = db.Model(&News{}).Count(&f.Total).Error
	}
	return
}

func (f *NewsFilter) GetListFrond(db *gorm.DB) (listNews []News,err error) {
	return
}

func(f NewsFilter) GetA(db *gorm.DB) (news News,err error) {
	err = db.First(&news,f.Id).Error
	return
}

func(f NewsFilter) Delete(db *gorm.DB) error {
	return db.Where("id = ?",f.Id).Delete(&News{}).Error
}

func (f NewsCategoryFilter) GetList(db *gorm.DB) (categories []NewsCategory,err error) {
	offset := f.getOffset()
	err = db.Limit(f.Rows).Offset(offset).Find(&categories).Error
	return
}

func(f NewsCategoryFilter) GetA(db *gorm.DB) (category NewsCategory,err error) {
	err = db.First(&category,f.Id).Error
	return
}

func(f NewsCategoryFilter) Delete(db *gorm.DB) error {
	return db.Where("id = ?",f.Id).Delete(&NewsCategory{}).Error
}