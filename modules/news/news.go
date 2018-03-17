package news

import (
	"gopkg.in/kataras/iris.v8"
	"github.com/nobita0590/web_mysql/models"
	"github.com/nobita0590/web_mysql/modules/common"
	"fmt"
)

func listNews(ctx iris.Context) {
	filter := models.NewsFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if listNews,err := filter.GetList(db);err == nil{
			common.Success(ctx,iris.Map{
				"models": listNews,
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
func getRoot(ctx iris.Context)  {
	filter1 := models.NewsFilter{
		CategoriesId: []uint{1},
		Page: models.Page{
			Rows: 5,
		},
	}
	filter2 := models.NewsFilter{
		CategoriesId: []uint{2},
		Page: models.Page{
			Rows: 5,
		},
	}
	db,_ := common.GetDb(ctx)
	listNews1,_ := filter1.GetList(db)
	listNews2,_ := filter2.GetList(db)
	res := make(iris.Map)
	if len(listNews1) > 0 {
		res["main1"] = listNews1[0]
		if len(listNews1) > 1 {
			res["news1"] = listNews1[1:]
		}
	}
	if len(listNews2) > 0 {
		res["main2"] = listNews2[0]
		if len(listNews2) > 1 {
			res["news2"] = listNews2[1:]
		}
	}
	common.Success(ctx,res)
}
func aNews(ctx iris.Context) {
	filter := models.NewsFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if news,err := filter.GetA(db);err == nil{
			// db.Exec("UPDATE news SET views = views + 1 WHERE id = ?",news.ID)
			common.Success(ctx,news)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}

func publicNews(ctx iris.Context) {
	filter := models.NewsFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if news,err := filter.GetA(db);err == nil{
			news.Views += 1
			fmt.Print(news.Views)
			//u := db.Model(&news).Update("views").Error
			db.Exec("UPDATE news SET views = views + 1 WHERE id = ?",news.ID)
			listFilter := models.NewsFilter{
				IgnoreIds: []uint{news.ID},
				CategoriesId: []uint{news.CategoryId},
				Page: filter.Page,
			}
			lDb := db.New()
			defer lDb.Close()
			listNews,_ := listFilter.GetList(lDb)
			common.Success(ctx,iris.Map{
				"news": news,
				"relate" : listNews,
			})
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}

func insertNews(ctx iris.Context) {
	db,_ := common.GetDb(ctx)
	var news models.News
	if err,validErrs := common.ReadJSONStruct(ctx,&news,true);err == nil{
		user,_ := common.GetUser(ctx)
		news.CreatorId = user.ID
		if err := db.Create(&news).Error;err == nil{
			common.Success(ctx,news.ID)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,validErrs)
	}
}

func updateNews(ctx iris.Context) {
	db,_ := common.GetDb(ctx)
	var news models.News
	if err,validErrs := common.ReadJSONStruct(ctx,&news,true);err == nil{
		news.PrettyUrl = ""
		if err := db.Model(&news).Updates(news).Error;err == nil{
			common.Success(ctx,news.ID)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,validErrs)
	}
}
func deleteNews(ctx iris.Context) {
	filter := models.NewsFilter{}
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
