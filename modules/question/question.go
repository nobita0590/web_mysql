package question

import (
	"gopkg.in/kataras/iris.v8"
	"github.com/nobita0590/web_mysql/models"
	"github.com/nobita0590/web_mysql/modules/common"
	"encoding/json"
)

func getListQuestion(ctx iris.Context) {
	filter := models.QuestionFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if listQuest,err := filter.GetList(db);err == nil{
			common.Success(ctx,
				iris.Map{
					"models": listQuest,
					"p_info": filter.GetPageInfo(),
				})
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}

func getAQuestion(ctx iris.Context) {
	filter := models.QuestionFilter{}
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

func insertQuestion(ctx iris.Context) {
	db,_ := common.GetDb(ctx)
	var question models.Question
	if err,validErrs := common.ReadJSONStruct(ctx,&question,true);err == nil{
		user,_ := common.GetUser(ctx)
		question.CreatorId = user.ID
		answerConvert,_ := json.Marshal(question.AnswerView)
		question.Answer = string(answerConvert)
		if err := db.Create(&question).Error;err == nil{
			common.Success(ctx,question.ID)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,validErrs)
	}
}

func updateQuestion(ctx iris.Context) {
	db,_ := common.GetDb(ctx)
	var question models.Question
	if err,validErrs := common.ReadJSONStruct(ctx,&question,true);err == nil{
		answerConvert,_ := json.Marshal(question.AnswerView)
		question.Answer = string(answerConvert)
		if err := db.Model(&question).Updates(question).Error;err == nil{
			common.Success(ctx,question.ID)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,validErrs)
	}
}
func deleteQuestion(ctx iris.Context) {
	filter := models.QuestionFilter{}
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
