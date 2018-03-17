package models

import (
	"github.com/jinzhu/gorm"
	"time"
	"encoding/json"
)

type (
	Course struct {
		gorm.Model
		Title		string		`gorm:"type:varchar(255)" validate:"required"`
		PrettyUrl	string		`gorm:"type:varchar(255);unique_index" validate:"required"`
		BoughtNumber	uint
		ImageUrl	string		``
		YoutubeUrl	string
		TeacherId		string		`validate:"required"`
		TeacherName		string		`gorm:"-"`
		Price		string
		IsSaleOff	bool
		SaleOffPrice	string
		SaleOffDescription		string
		Tags		string

		Description	string		`validate:"required"`
		Content		string		`json:"-"`
		ContentDetail	[]CourseChapter		`json:"Content"`
		Benefit		string		`validate:"required"`
		Target		string		`validate:"required"`
		Interest	string		`validate:"required"`
		Resitered	int			`gorm:"-"`
		StudentsNumber		int		``

		StartDate	time.Time
		EndDate		time.Time

		CreatorId	uint
		CreatorName		string		`gorm:"-" json:",omitempty" validate:"-"`
		Registers		[]CourseRegister		`json:"-" form:"-"`
	}
	CourseChapter	struct {
		Title		string
		Steps		[]CourseStep
	}
	CourseStep	struct {
		Name		string
	}
	CourseFilter struct {
		Page
		Id		uint		`form:"ID"`
		IgnoreIds  		[]uint
	}
)

func (f *CourseFilter) GetList(db *gorm.DB) (courses []Course,err error) {
	db = db.Select("courses.*,t.value as teacher_name").
		Joins("LEFT JOIN select_sources t ON t.id = courses.teacher_id")
	offset := f.getOffset()
	if len(f.IgnoreIds) > 0 {
		db = db.Where("courses.id NOT IN (?)",f.IgnoreIds)
	}
	err = db.Limit(f.Rows).Offset(offset).Find(&courses).Error
	if f.Count {
		err = db.Model(&Course{}).Count(&f.Total).Error
	}
	return
}

func (f *CourseFilter) GetListFrond(db *gorm.DB) (courses []Course,err error) {
	return
}

func(f CourseFilter) GetA(db *gorm.DB) (course Course,err error) {
	db = db.Select("courses.*,t.value as teacher_name").
			Joins("LEFT JOIN select_sources t ON t.id = courses.teacher_id")
	err = db.First(&course,f.Id).Error
	if err == nil {
		err = json.Unmarshal([]byte(course.Content),&course.ContentDetail)
	}
	return
}

func(f CourseFilter) Delete(db *gorm.DB) error {
	return db.Where("id = ?",f.Id).Delete(&Course{}).Error
}

/*
func (c *Course) CheckUser(userId uint) ([]string,bool) {
	usersId := strings.Split(c.UsersId,",")
	userStrId := strconv.Itoa(int(userId))
	for _,id := range usersId {
		if userStrId == id {
			c.Resitered = 1
			return usersId,true
		}
	}
	return usersId,false
}*/
