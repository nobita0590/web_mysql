package models

import (
	"github.com/jinzhu/gorm"
)

//import "github.com/jinzhu/gorm"

type (
	SelectSource struct {
		ID        	uint 		`gorm:"primary_key"`
		Value		string		`gorm:"varchar(255)"`
		ConvertedValue		string		`gorm:"varchar(255)"`
		GroupId		uint
		Extra 		string
		IsSystem	uint
		Order		uint
		RelateId	uint
	}
	SelectSourceFilter struct {
		Page
		Id 			uint		`form:"ID"`
		Ids 		[]uint
		GroupsId	[]uint
	}
)

func (f *SelectSourceFilter) GetList(db *gorm.DB) (selects []SelectSource,err error) {
	// offset := f.getOffset()
	if f.Sort == "" {
		f.Sort = "order"
	}
	db = db.Order(f.getOrderBy())
	if len(f.GroupsId) > 0 {
		db = db.Where("group_id in (?)",f.GroupsId)
	}
	err = db.Find(&selects).Error
	if f.Count && err == nil {
		err = db.Model(&User{}).Count(&f.Total).Error
	}
	if len(f.Ids) > 0 {
		db = db.Where("id in (?)",f.Ids)
	}
	return
}

func(f SelectSourceFilter) GetA(db *gorm.DB) (item SelectSource,err error) {
	err = db.First(&item,f.Id).Error
	return
}

func(f SelectSourceFilter) Delete(db *gorm.DB) error {
	return db.Where("id = ?",f.Id).Where("is_system = 0").Delete(&SelectSource{}).Error
}