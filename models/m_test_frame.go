package models

import (
	"github.com/jinzhu/gorm"
	"encoding/json"
)

type (
	TestsFrame struct {
		gorm.Model
		Title		string		`gorm:"type:varchar(255)" validate:"required"`
		SubjectId	uint
		ClassesId	uint
		TypeId		uint
		TypeName	string		`form:"-" gorm:"-"`
		SubjectName	string		`form:"-" gorm:"-"`
		ClassesName	string		`form:"-" gorm:"-"`
		CreatorName	string		`form:"-" gorm:"-"`
		DifficulConfig		string				`form:"-" json:"-"`
		DifficulConfigView	[]DifficulConfig	`gorm:"-" json:"DifficulConfig" form:"DifficulConfig"`
		Description		string
		Total		uint
		Time 		uint
		CreatorId	uint
		Minutes		int
	}
	DifficulConfig	struct {
		DifficulId	uint
		Percent		uint
	}
	TestsFrameFilter struct {
		Page
		Id		uint		`form:"ID"`
		SubjectId 		uint
		ClassesId 		uint
	}
)

func (f *TestsFrameFilter) GetList(db *gorm.DB) (testsFrames []TestsFrame,err error) {
	offset := f.getOffset()
	/*err = db.Table("tests_frames").
		Select("tests_frames.*,t_type.value type_name,class.value class_name,subject.value subject_name").
		Joins("left join select_sources t_type on tests_frames.type_id = t_type.id").
		Joins("left join select_sources class on tests_frames.classes_id = class.id").
		Joins("left join select_sources subject on tests_frames.subject_id = subject.id").
		Where("tests_frames.deleted_at IS NULL").*/
	err = db.Table("tests_frames").
		Select("tests_frames.*,type.value type_name,type.extra minutes" +
			",CONCAT(users.last_name,' ',users.first_name) creator_name").
		Joins("left join select_sources type on tests_frames.type_id = type.id").
		Joins("left join users on users.id = tests_frames.creator_id").
		//Where("tests.deleted_at IS NULL").
		Limit(f.Rows).Offset(offset).Find(&testsFrames).Error
	// err = db.Model(&TestsFrame{}).Limit(f.Rows).Offset(offset).Find(&testsFrames).Error
	if f.Count && err == nil {
		err = db.Model(&TestsFrame{}).Count(&f.Total).Error
	}
	return
}

func(f TestsFrameFilter) GetA(db *gorm.DB) (test TestsFrame,err error) {
	err = db.First(&test,f.Id).Error
	if err == nil && len(test.DifficulConfig) > 0 {
		err = json.Unmarshal([]byte(test.DifficulConfig), &test.DifficulConfigView)
	}
	return
}

func(f TestsFrameFilter) Delete(db *gorm.DB) error {
	return db.Where("id = ?",f.Id).Delete(&TestsFrame{}).Error
}

func(t TestsFrame) GetConfig() map[uint]uint {
	config := map[uint]uint{}
	var total,filled uint = 0,0
	for _,v := range t.DifficulConfigView {
		total += v.Percent
	}
	lenConf := len(t.DifficulConfigView)
	for i:= 0;i < len(t.DifficulConfigView);i++ {
		var cancu uint
		if (i + 1) < lenConf {
			cancu = t.Total * t.DifficulConfigView[i].Percent / total
			filled += cancu
			config[t.DifficulConfigView[i].DifficulId] = cancu//t.Total * t.DifficulConfigView[i].Percent / total
		}else{
			cancu = t.Total - filled
			config[t.DifficulConfigView[i].DifficulId] = cancu//total - filled
		}
	}
	return config
}