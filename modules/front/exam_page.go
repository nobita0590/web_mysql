package front

import (
	"gopkg.in/kataras/iris.v8"
	"github.com/nobita0590/web_mysql/modules/common"
	"encoding/json"
	"github.com/nobita0590/web_mysql/models"
)

func examInsert(ctx iris.Context)  {
	db,_ := common.GetDb(ctx)
	var exams models.Exams
	if err,validErrs := common.ReadStruct(ctx,&exams,true);err == nil{
		user,_ := common.GetUser(ctx)
		exams.UserId = user.ID
		content,_ := json.Marshal(exams.HistoryDetail)
		exams.History = string(content)
		testTable := "tests_frames"
		if (exams.TypeId != 2) {
			exams.TypeId = 1;
			testTable = "tests"
		}
		if err := db.Create(&exams).Error;err == nil{
			db.Exec("UPDATE "+testTable+" SET time = time + 1 WHERE id = ?",exams.TestId)
			common.Success(ctx,exams.ID)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,validErrs)
	}
}