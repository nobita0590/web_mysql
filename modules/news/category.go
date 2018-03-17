package news

import (
	"gopkg.in/kataras/iris.v8"
	"github.com/nobita0590/web_mysql/models"
	"github.com/nobita0590/web_mysql/modules/common"
)

func listCategory(ctx iris.Context) {
	filter := models.NewsCategoryFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if categories,err := filter.GetList(db);err == nil{
			common.Success(ctx,categories)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}

func aCategory(ctx iris.Context) {
	filter := models.NewsCategoryFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if category,err := filter.GetA(db);err == nil{
			common.Success(ctx,category)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}

func insertCategory(ctx iris.Context) {
	db,_ := common.GetDb(ctx)
	var category models.NewsCategory
	if err,validErrs := common.ReadStruct(ctx,&category,true);err == nil{
		if err := db.Create(&category).Error;err == nil{
			common.Success(ctx,category.ID)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,validErrs)
	}
}

func updateCategory(ctx iris.Context) {
	db,_ := common.GetDb(ctx)
	var category models.NewsCategory
	if err,validErrs := common.ReadStruct(ctx,&category,true);err == nil{
		if err := db.Model(&category).Updates(category).Error;err == nil{
			common.Success(ctx,category.ID)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,validErrs)
	}
}
func deleteCategory(ctx iris.Context) {

}
