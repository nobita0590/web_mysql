package question

import (
	"gopkg.in/kataras/iris.v8"
	"github.com/nobita0590/web_mysql/models"
	"github.com/nobita0590/web_mysql/modules/common"
	"strings"
)

func getListTests(ctx iris.Context) {
	filter := models.TestsFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		user,_ := common.GetUser(ctx)
		if listTests,err := filter.GetList(db,user.ID);err == nil{
			common.Success(ctx,
				iris.Map{
					"models": listTests,
					"p_info": filter.GetPageInfo(),
				})
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}
func ListTestForFrontend(ctx iris.Context) {
	filter := models.TestsFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if listTests,err := filter.GetListForFront(db);err == nil{
			common.Success(ctx,
				iris.Map{
					"models": listTests,
					"p_info": filter.GetPageInfo(),
				})
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}

func getATests(ctx iris.Context) {
	filter := models.TestsFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if quest,err := filter.GetA(db);err == nil{
			common.Success(ctx,quest)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}
func getTestsExams(ctx iris.Context)  {
	filter := models.TestsFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if test,err := filter.GetA(db);err == nil{
			examsFilter := models.ExamsFilter{}
			exams,err := examsFilter.ChartByTests(db,test.ID)
			common.Success(ctx,iris.Map{
				"test" : test,
				"exams" : exams,
				"err" : err,
			})
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}
func getTopTen(ctx iris.Context) {
	filter := models.ExamsFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if tops,err := filter.ChartTopTen(db);err == nil {
			common.Success(ctx,tops)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}
func insertTests(ctx iris.Context) {
	db,_ := common.GetDb(ctx)
	var tests models.Tests
	if err,validErrs := common.ReadStruct(ctx,&tests,true);err == nil{
		user,_ := common.GetUser(ctx)
		tests.CreatorId = user.ID
		tests.QuestionsId = strings.Join(tests.QuestsId,",")
		if err := db.Create(&tests).Error;err == nil{
			common.Success(ctx,tests.ID)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,validErrs)
	}
}

func updateTests(ctx iris.Context) {
	db,_ := common.GetDb(ctx)
	var tests models.Tests
	if err,validErrs := common.ReadStruct(ctx,&tests,true);err == nil{
		tests.QuestionsId = strings.Join(tests.QuestsId,",")
		if err := db.Model(&tests).Updates(tests).Error;err == nil{
			common.Success(ctx,tests.ID)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,validErrs)
	}
}
func deleteTests(ctx iris.Context) {
	filter := models.TestsFilter{}
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
