package front

import (
	"gopkg.in/kataras/iris.v8"
	"github.com/nobita0590/web_mysql/models"
	"github.com/nobita0590/web_mysql/modules/common"
)
func listFags(ctx iris.Context) {
	filter := models.FagsFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		user,_ := common.GetUser(ctx)
		if fags,err := filter.GetList(db,user.ID);err == nil{
			common.Success(ctx,iris.Map{
				"models": fags,
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

func aFags(ctx iris.Context) {
	filter := models.FagsFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if fag,err := filter.GetA(db);err == nil{
			common.Success(ctx,fag)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}

func insertFags(ctx iris.Context) {
	db,_ := common.GetDb(ctx)
	var fag models.Fags
	if err,validErrs := common.ReadJSONStruct(ctx,&fag,true);err == nil{
		user,_ := common.GetUser(ctx)
		fag.UserId = user.ID
		if err := db.Create(&fag).Error;err == nil{
			common.Success(ctx,fag)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,validErrs)
	}
}

func updateFags(ctx iris.Context) {
	db,_ := common.GetDb(ctx)
	var fag models.Fags
	if err,validErrs := common.ReadJSONStruct(ctx,&fag,true);err == nil{
		fag.PrettyUrl = ""
		if err := db.Model(&fag).Updates(fag).Error;err == nil{
			common.Success(ctx,fag)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,validErrs)
	}
}
func deleteFags(ctx iris.Context) {
	filter := models.FagsFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if err := filter.Delete(db);err == nil{
			common.Success(ctx,filter.ID)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}
