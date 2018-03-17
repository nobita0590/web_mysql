package question

import (
	"gopkg.in/kataras/iris.v8"
	"github.com/nobita0590/web_mysql/models"
	"github.com/nobita0590/web_mysql/modules/common"
	"encoding/json"
)

func getListTestsFrame(ctx iris.Context) {
	filter := models.TestsFrameFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if listTestsFrame,err := filter.GetList(db);err == nil{
			common.Success(ctx,
				iris.Map{
					"models": listTestsFrame,
					"p_info": filter.GetPageInfo(),
				})
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}

func getATestsFrame(ctx iris.Context) {
	filter := models.TestsFrameFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if testFrame,err := filter.GetA(db);err == nil{
			common.Success(ctx,testFrame)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}

func insertTestsFrame(ctx iris.Context) {
	db,_ := common.GetDb(ctx)
	var testsFrame models.TestsFrame
	if err,validErrs := common.ReadStruct(ctx,&testsFrame,true);err == nil{
		user,_ := common.GetUser(ctx)
		testsFrame.CreatorId = user.ID
		configConvert,_ := json.Marshal(testsFrame.DifficulConfigView)
		testsFrame.DifficulConfig = string(configConvert)
		if err := db.Create(&testsFrame).Error;err == nil{
			common.Success(ctx,testsFrame.ID)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,validErrs)
	}
}

func updateTestsFrame(ctx iris.Context) {
	db,_ := common.GetDb(ctx)
	var testsFrame models.TestsFrame
	if err,validErrs := common.ReadStruct(ctx,&testsFrame,true);err == nil{
		configConvert,_ := json.Marshal(testsFrame.DifficulConfigView)
		testsFrame.DifficulConfig = string(configConvert)
		if err := db.Model(&testsFrame).Updates(testsFrame).Error;err == nil{
			common.Success(ctx,testsFrame.ID)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,validErrs)
	}
}
func deleteTestsFrame(ctx iris.Context) {
	filter := models.TestsFrameFilter{}
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

func generateExamFrame(ctx iris.Context) {
	filter := models.TestsFrameFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if frame,err := filter.GetA(db);err == nil {
			config := frame.GetConfig()
			qFilter := models.QuestionFilter{
				ClassId: []uint{filter.ClassesId},
				SubjectId: []uint{filter.SubjectId},
			}
			questions := []models.Question{}
			for difficulId,getNumber := range config {
				qFilter.DifficultId = []uint{difficulId}
				qFilter.Rows = int(getNumber)
				sliceQuests,_ := qFilter.GetRandom(db.New())
				questions = append(questions,sliceQuests...)
			}
			common.Success(ctx,questions)
		}else{
			common.BadRequest(ctx,err,"Khong tim thay loai de thi")
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}
