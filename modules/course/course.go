package course

import (
	"gopkg.in/kataras/iris.v8"
	"github.com/nobita0590/web_mysql/models"
	"github.com/nobita0590/web_mysql/modules/common"
	"encoding/json"
	"errors"
	"time"
	"fmt"
)

func listCourse(ctx iris.Context) {
	filter := models.CourseFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if courses,err := filter.GetList(db);err == nil{
			common.Success(ctx,iris.Map{
				"models": courses,
				"p_info": filter.GetPageInfo(),
			})
			//common.Success(ctx,listNews)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}

func aCourse(ctx iris.Context) {
	filter := models.CourseFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if course,err := filter.GetA(db);err == nil{
			common.Success(ctx,course)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}

func insertCourse(ctx iris.Context) {
	db,_ := common.GetDb(ctx)
	var course models.Course
	if err,validErrs := common.ReadJSONStruct(ctx,&course,true);err == nil{
		user,_ := common.GetUser(ctx)
		course.CreatorId = user.ID
		content,_ := json.Marshal(course.ContentDetail)
		course.Content = string(content)
		if err := db.Create(&course).Error;err == nil{
			common.Success(ctx,course.ID)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,validErrs)
	}
}

func updateCourse(ctx iris.Context) {
	db,_ := common.GetDb(ctx)
	var course models.Course
	if err,validErrs := common.ReadJSONStruct(ctx,&course,true);err == nil{
		content,_ := json.Marshal(course.ContentDetail)
		course.Content = string(content)
		if err := db.Model(&course).Updates(course).Error;err == nil{
			common.Success(ctx,course.ID)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,validErrs)
	}
}
func deleteCourse(ctx iris.Context) {
	filter := models.CourseFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if err := filter.Delete(db);err == nil{
			common.Success(ctx,filter.Id)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}
func publicCourse(ctx iris.Context) {
	filter := models.CourseFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if course,err := filter.GetA(db);err == nil{
			if user,err := common.GetUser(ctx);err == nil{
				register := models.CourseRegister{}
				if err := db.Where("course_id = ?", course.ID).
					Where("user_id = ?",user.ID).First(&register).Error;err == nil {
					course.Resitered = 1
				}
			}
			listFilter := models.CourseFilter{
				IgnoreIds: []uint{course.ID},
				// CategoriesId: []uint{news.CategoryId},
				Page: filter.Page,
			}
			lDb := db.New()
			defer lDb.Close()
			courses,_ := listFilter.GetList(lDb)
			common.Success(ctx,iris.Map{
				"course": course,
				"relate" : courses,
			})
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}

func RegisCourse(ctx iris.Context)  {
	filter := models.CourseFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		fmt.Printf(`%+v`,filter)
		if course,err := filter.GetA(db);err == nil{
			user,_ := common.GetUser(ctx)
			register := models.CourseRegister{}
			db = db.New()
			if err := db.Where("course_id = ?", course.ID).
				Where("user_id = ?",user.ID).First(&register).Error;err == nil {
				common.BadRequest(ctx,errors.New("Bạn đã đăng ký khóa học này rồi"),nil)
			}else{
				register = models.CourseRegister{
					UserId:user.ID,
					CourseId:course.ID,
					CreatedAt:time.Now(),
				}
				db.Create(&register)
				course.Resitered = 1
				course.StudentsNumber += 1
				db.Model(&course).Updates(map[string]interface{}{"students_number":course.StudentsNumber})
				common.Success(ctx,course)
			}
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}

func listRegisCourse(ctx iris.Context) {
	filter := models.CourseRegisterFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if courseRegisters,err := filter.GetList(db);err == nil{
			common.Success(ctx,iris.Map{
				"models": courseRegisters,
				"p_info": filter.GetPageInfo(),
			})
			//common.Success(ctx,listNews)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}