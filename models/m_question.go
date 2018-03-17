package models

import (
	"github.com/jinzhu/gorm"
	"encoding/json"
)

type(
	Question struct {
		gorm.Model
		Title 		string
		Content		string
		FullAnswer		string
		CategoryId		uint
		ClassId			uint
		DifficultId		uint
		SubCategoryId	uint
		SubjectId		uint
		CreatorId		uint					`form:"-"`
		Answer			string					`json:"-"`
		AnswerView		[]AnswerOption			`gorm:"-" validate:"min=2"`

		CategoryName	string					`gorm:"-"`
		SubCategoryName	string					`gorm:"-"`
		ClassName		string					`gorm:"-"`
		DifficultName	string					`gorm:"-"`
		SubjectName		string					`gorm:"-"`
		CreatorName		string					`gorm:"-"`
	}
	AnswerOption struct {
		No 			uint
		IsTrue		bool
		Content		string
	}
	QuestionFilter struct {
		Page
		Id			uint		`form:"ID"`
		Ids 		[]string
		ClassId  	[]uint
		SubjectId  	[]uint
		CategoryId  []uint
		DifficultId []uint
	}
)

func (f *QuestionFilter) BindWhere(db *gorm.DB) *gorm.DB {
	if len(f.ClassId) > 0 {
		db = db.Where("questions.class_id IN (?)",f.ClassId)
	}
	if len(f.SubjectId) > 0 {
		db = db.Where("questions.subject_id IN (?)",f.SubjectId)
	}
	if len(f.CategoryId) > 0 {
		db = db.Where("questions.category_id IN (?)",f.CategoryId)
	}
	if len(f.DifficultId) > 0 {
		db = db.Where("questions.difficult_id IN (?)",f.DifficultId)
	}
	if len(f.Ids) > 0 {
		db = db.Where("questions.id IN (?)",f.Ids)
	}
	return db
}

func (f *QuestionFilter) GetList(db *gorm.DB) (questions []Question,err error) {
	offset := f.getOffset()
	db = f.BindWhere(db)
	rows,err := db.Table("questions").
		Select("questions.*,CONCAT(users.last_name,' ',users.first_name) creator_name,sub_categories.value sub_category_name").
		Joins("left join users on users.id = questions.creator_id").//SubCategoryName
		Joins("left join sub_categories on sub_categories.id = questions.sub_category_id").
		Where("questions.deleted_at IS NUll").
		Limit(f.Rows).Offset(offset).Rows()
	if err != nil {
		return
	}
	question := Question{}
	selectsFilter := SelectSourceFilter{}
	for rows.Next() {
		question = Question{}
		err = db.ScanRows(rows, &question)
		if err == nil {
			err = json.Unmarshal([]byte(question.Answer),&question.AnswerView)
			if err == nil {
				questions = append(questions,question)
			}
			selectsFilter.Ids = append(selectsFilter.Ids,
				question.ClassId, question.CategoryId, question.DifficultId, question.SubjectId)
		}
	}
	rows.Close()
	if selectsSource,err := selectsFilter.GetList(db.New());err == nil {
		for k,question := range questions{
			question.FillSource(selectsSource)
			questions[k] = question
		}
	}
	if f.Count{
		//db = f.BindWhere(db)
		err = db.Model(&Question{}).Count(&f.Total).Error
	}
	// err = db.Limit(f.Rows).Offset(offset).Find(&documents).Error
	return
}

func (q *Question) FillSource(selectsSource []SelectSource){
	for _,source := range selectsSource {
		if q.SubjectId == source.ID {
			q.SubjectName = source.Value
		}
		if q.CategoryId == source.ID {
			q.CategoryName = source.Value
		}
		if q.ClassId == source.ID {
			q.ClassName = source.Value
		}
		if q.DifficultId == source.ID {
			q.DifficultName = source.Value
		}
	}
}

func(f QuestionFilter) GetA(db *gorm.DB) (question Question,err error) {
	err = db.First(&question,f.Id).Error
	if err == nil {
		err = json.Unmarshal([]byte(question.Answer),&question.AnswerView)
	}
	return
}

func(f QuestionFilter) Delete(db *gorm.DB) error {
	return db.Where("id = ?",f.Id).Delete(&Question{}).Error
}


func (f *QuestionFilter) GetRandom(db *gorm.DB) (questions []Question,err error) {
	rows,err := db.Table("questions").Select("*").
		Where("deleted_at IS NUll").
		Where("subject_id in (?)",f.SubjectId).
		Where("difficult_id in (?)",f.DifficultId).
		Where("class_id in (?)",f.ClassId).
		Order("rand()").
		Limit(f.Rows).Rows()
	if err != nil {
		return
	}
	question := Question{}
	selectsFilter := SelectSourceFilter{}
	for rows.Next() {
		question = Question{}
		err = db.ScanRows(rows, &question)
		if err == nil {
			err = json.Unmarshal([]byte(question.Answer),&question.AnswerView)
			if err == nil {
				questions = append(questions,question)
			}
			selectsFilter.Ids = append(selectsFilter.Ids,
				question.ClassId, question.CategoryId, question.DifficultId, question.SubjectId)
		}
	}
	rows.Close()
	if selectsSource,err := selectsFilter.GetList(db.New());err == nil {
		for k,question := range questions{
			question.FillSource(selectsSource)
			questions[k] = question
		}
	}
	return
}