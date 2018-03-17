package models

import "github.com/jinzhu/gorm"

type (
	SubCategory struct {
		ID        	uint 		`gorm:"primary_key"`
		Value		string		`gorm:"varchar(255)"`
		ConvertedValue		string		`gorm:"varchar(255)"`
		GroupId		uint
		Extra 		string
		IsSystem	uint
		Order		uint
		RelateId	uint
	}
	SubCategoryFilter struct {
		Page
		Id 			uint		`form:"ID"`
		Ids 		[]uint
		GroupsId	[]uint
	}
)

func (SubCategory) TableName() string {
	return "sub_categories"
}

func (f *SubCategoryFilter) GetList(db *gorm.DB) (selects []SubCategory,err error) {
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

func(f SubCategoryFilter) GetA(db *gorm.DB) (item SubCategory,err error) {
	err = db.First(&item,f.Id).Error
	return
}

func(f SubCategoryFilter) Delete(db *gorm.DB) error {
	return db.Where("id = ?",f.Id).Delete(&SubCategory{}).Error
}