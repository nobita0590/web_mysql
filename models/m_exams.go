package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

/*
export class Exams {

}
*/
type (
	Exams struct {
		ID 			uint
		UserId 		uint
		TestId 		uint
		TypeId  	uint
		UserName	string		`gorm:"-"`
		TestName	string		`gorm:"-"`
		Total 		int
		TrueNumber 	int
		Score 		float32
		StartTime 	time.Time
		FinishTime 	time.Time
		TimeDoing 	int
		History 	string		`json:"-"`
		HistoryDetail	[]QuestionHistory		`json:"History" gorm:"-" form:"History"`
	}
	ExamsFilter struct {
		Page
		IDs 		[]uint
		TestIds		[]uint
		TypeId		uint
		From		time.Time
		To			time.Time
	}
	QuestionHistory struct {
		True		uint
		Picked		uint
		QuestionId	uint
	}
	TopTenUser struct {
		ID 			uint
		TScore		float64
		AvgScore	float64
		UserName	string
	}
)

func (f *ExamsFilter) ChartByTests(db *gorm.DB,testId uint) (listExams []Exams,err error) {
	if f.TypeId != 2 {
		f.TypeId = 1
	}
	err = db.Table("exams").
		Select("exams.total,exams.true_number,exams.score,exams.time_doing,CONCAT(users.last_name,' ',users.first_name) user_name").
		Joins("left join users on users.id = exams.user_id").
		Where("exams.type_id = ?", f.TypeId).
		Where("exams.test_id = ?", testId).
		Order("exams.score desc").
		Order("exams.time_doing").
		Limit(10).Find(&listExams).Error
	return
}
func (f *ExamsFilter) ChartTopTen(db *gorm.DB) (tops []TopTenUser,err error) {
	var from,to = f.From.UnixNano(),f.To.UnixNano()
	if from > 0 && to > from {
		db = db.Where("exams.start_time BETWEEN ? AND ?",f.From,f.To)
	}else{
		if from > 0 {
			db = db.Where("exams.start_time > ?",f.From)
		}
		if to > 0 {
			db = db.Where("exams.start_time < ?",f.To)
		}
	}

	err = db.Table("exams").
		Select("users.id,sum(score) t_score,avg(score) avg_score,CONCAT(users.last_name,' ',users.first_name) user_name").
		Joins("LEFT JOIN users on users.id = exams.user_id").
		Group("exams.user_id").
		Order("avg_score desc").
		Order("t_score desc").
		Limit(10).Find(&tops).Error
	return
}