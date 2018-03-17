package models

import (
	"github.com/jinzhu/gorm"
	"strings"
)

type (
	Tests struct {
		gorm.Model
		Title		string		`gorm:"type:varchar(255)" validate:"required"`
		PrettyUrl	string		`gorm:"type:varchar(255);unique_index" `
		Description	string		`gorm:"type:varchar(1023)"`
		TypeId		uint
		Time		uint
		CreatorId	uint
		ClassId 	uint
		SubjectId 	uint
		TypeName	string		`gorm:"-" form:"-"`
		ClassName	string		`gorm:"-" form:"-"`
		SubjectName	string		`gorm:"-" form:"-"`
		CreatorName	string		`gorm:"-" form:"-"`
		QuestionsId	string		`form:"-" json:"-"`
		QuestsId 	[]string	`gorm:"-" form:"QuestionsId" json:"QuestionsId,omitempty"`
		Questions 	[]Question	`gorm:"-" form:"-" json:"Questions,omitempty"`
		Minutes		int			`gorm:"-" form:"-"`
	}
	TestsFilter struct {
		Page
		Id		uint		`form:"ID"`
		UserFilterAction	string	`form:"uf"`
		ClassId			uint
		SubjectId		uint
		TypeId			uint
	}
)

func (f *TestsFilter) GetList(db *gorm.DB,userId uint) (listTests []Tests,err error) {
	offset := f.getOffset()
	if f.ClassId > 0 {
		db = db.Where("tests.class_id = ?",f.ClassId)
	}
	if f.SubjectId > 0 {
		db = db.Where("tests.subject_id = ?",f.SubjectId)
	}
	if f.TypeId > 0 {
		db = db.Where("tests.type_id = ?",f.TypeId)
	}
	switch f.UserFilterAction {
	case "doing":
	case "done":
		db = db.Joins("left join exams on exams.test_id = tests.id").
			Where("exams.type_id = 1").
			Where("exams.user_id = ?",userId).
			Group("tests.id")
	}
	err = db.Table("tests").
		Select("tests.*,t_type.value type_name,class.value class_name,subject.value subject_name,CONCAT(users.last_name,' ',users.first_name) creator_name").
		Joins("left join select_sources t_type on tests.type_id = t_type.id").
		Joins("left join select_sources class on tests.class_id = class.id").
		Joins("left join select_sources subject on tests.subject_id = subject.id").
		Joins("left join users on users.id = tests.creator_id").
		Where("tests.deleted_at IS NULL").
		Limit(f.Rows).Offset(offset).Find(&listTests).Error
	//err = db.Model(&News{}).Limit(f.Rows).Offset(offset).Find(&listTests).Error
	if f.Count && err == nil {
		err = db.Model(&Tests{}).Count(&f.Total).Error
	}
	return
}
func (f *TestsFilter) GetListForFront(db *gorm.DB) (listTests []Tests,err error) {
	offset := f.getOffset()
	err = db.Table("tests").
		Select("tests.*,t_type.value type_name,class.value class_name,subject.value subject_name").
		Joins("left join select_sources t_type on tests.type_id = t_type.id").
		Joins("left join select_sources class on tests.class_id = class.id").
		Joins("left join select_sources subject on tests.subject_id = subject.id").
		Where("tests.deleted_at IS NULL").
		Limit(f.Rows).Offset(offset).Find(&listTests).Error
	//err = db.Model(&News{}).Limit(f.Rows).Offset(offset).Find(&listTests).Error
	if f.Count && err == nil {
		err = db.Model(&Tests{}).Count(&f.Total).Error
	}
	return
}

func(f TestsFilter) GetA(db *gorm.DB) (test Tests,err error) {
	err = db.Table("tests").
		Select("tests.*,t_type.value type_name,t_type.extra minutes,class.value class_name,subject.value subject_name").
		Joins("left join select_sources t_type on tests.type_id = t_type.id").
		Joins("left join select_sources class on tests.class_id = class.id").
		Joins("left join select_sources subject on tests.subject_id = subject.id").
		Where("tests.id = ?", f.Id).First(&test).Error
	// err = db.First(&test,f.Id).Error
	test.QuestsId = strings.Split(test.QuestionsId,",")
	if len(test.QuestsId) > 0 {
		qFilter := QuestionFilter{
			Ids: test.QuestsId,
			Page: Page{Rows:999},
		}
		qDb := db.New()
		test.Questions, _ = qFilter.GetList(qDb)
	}
	return
}

func(f TestsFilter) Delete(db *gorm.DB) error {
	return db.Where("id = ?",f.Id).Delete(&Tests{}).Error
}