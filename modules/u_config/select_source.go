package u_config

import (
	"gopkg.in/kataras/iris.v8"
	"github.com/nobita0590/web_mysql/models"
	"github.com/nobita0590/web_mysql/modules/common"
	"strconv"
	"strings"
	"fmt"
)

func getListSelect(ctx iris.Context) {
	filter := models.SelectSourceFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		ctx.ReadForm(&filter)
		if listSelect,err := filter.GetList(db);err == nil{
			common.Success(ctx,listSelect)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}
func getListAndGroupSelect(ctx iris.Context) {
	filter := models.SelectSourceFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		filter.Count = false
		if listSource,err := filter.GetList(db);err == nil{
			groups := make(map[uint][]models.SelectSource)
			for _,source := range listSource {
				if _,ok := groups[source.GroupId];ok {
					groups[source.GroupId] = append(groups[source.GroupId],source)
				}else{
					groups[source.GroupId] = []models.SelectSource{source}
				}
			}
			common.Success(ctx,groups)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}

func getASelect(ctx iris.Context) {
	filter := models.SelectSourceFilter{}
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

func insertSelect(ctx iris.Context) {
	db,_ := common.GetDb(ctx)
	var item models.SelectSource
	if err,validErrs := common.ReadStruct(ctx,&item,true);err == nil{
		if err := db.Create(&item).Error;err == nil{
			common.Success(ctx,item.ID)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,validErrs)
	}
}

func updateSelect(ctx iris.Context) {
	db,_ := common.GetDb(ctx)
	var item models.SelectSource
	if err,validErrs := common.ReadStruct(ctx,&item,true);err == nil{
		if err := db.Model(&item).Updates(item).Error;err == nil{
			common.Success(ctx,item.ID)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,validErrs)
	}
}
func deleteSelect(ctx iris.Context) {
	filter := models.SelectSourceFilter{}
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

func updateSelectOrder(ctx iris.Context) {
	sortVal := struct {
		Sort map[int]int	`form:"sort"`
	}{}
	if err,_ := common.ReadStruct(ctx, &sortVal, false);err == nil {
		fmt.Println(sortVal)
		db,_ := common.GetDb(ctx)
		sql := "UPDATE `select_sources` SET `order` = CASE"
		ids := []string{}
		for id,order := range sortVal.Sort {
			idStr := strconv.Itoa(id)
			sql += " WHEN id = "+idStr+" THEN "+strconv.Itoa(order)
			ids = append(ids,idStr)
		}
		sql += " ELSE `order` END WHERE id in ("+strings.Join(ids,",")+")"
		if err := db.Exec(sql).Error; err == nil {
			common.Success(ctx,true)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}
