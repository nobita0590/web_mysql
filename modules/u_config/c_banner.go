package u_config

import (
	"github.com/nobita0590/web_mysql/modules/common"
	"github.com/nobita0590/web_mysql/models"
	"gopkg.in/kataras/iris.v8"
	"encoding/json"
)

func getListBanner(ctx iris.Context) {
	filter := models.BannerFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if listBanner,err := filter.GetList(db);err == nil{
			common.Success(ctx,listBanner)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}

func getABanner(ctx iris.Context) {
	filter := models.BannerFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if banner,err := filter.GetA(db);err == nil{
			common.Success(ctx,banner)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}

func updateBanner(ctx iris.Context) {
	db,_ := common.GetDb(ctx)
	var item models.Banner
	if err,validErrs := common.ReadStruct(ctx,&item,true);err == nil{
		info,_ := json.Marshal(item.InfoView)
		item.Info = string(info)
		if err := db.Model(&item).Updates(item).Error;err == nil{
			common.Success(ctx,item.ID)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,validErrs)
	}
}