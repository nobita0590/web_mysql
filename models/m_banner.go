package models

import (
	"github.com/jinzhu/gorm"
	"encoding/json"
)

type (
	Banner struct {
		ID 		uint
		Name 			string
		Position		string
		Info			string		`json:"-" form:"-"`
		InfoView		[]bannerInfo	`gorm:"-" json:"Info" form:"Info"`
		Number 			uint
	}
	BannerFilter struct {
		Page
		ID 		uint
		Position 		string
	}
	bannerInfo struct {
		Image 		string
		Link 		string
	}
)

func (f *BannerFilter) GetList(db *gorm.DB) (banners []Banner,err error) {
	rows,err := db.Table("banners").
		Select("banners.*").Rows()
	if err != nil {
		return
	}
	banner := Banner{}
	for rows.Next() {
		banner = Banner{}
		err = db.ScanRows(rows, &banner)
		if err == nil {
			e := json.Unmarshal([]byte(banner.Info),&banner.InfoView)
			if e != nil {
				banner.InfoView = []bannerInfo{}
			}
			banners = append(banners,banner)
		}
	}
	rows.Close()
	return
}

func(f BannerFilter) GetA(db *gorm.DB) (banner Banner,err error) {
	err = db.First(&banner,f.ID).Error
	if err == nil {
		err = json.Unmarshal([]byte(banner.Info),&banner.InfoView)
	}
	return
}

