package front

import (
	"gopkg.in/kataras/iris.v8"
	"github.com/nobita0590/web_mysql/modules/common"
	"github.com/nobita0590/web_mysql/models"
)

func newsPage(ctx iris.Context) {
	db,_ := common.GetDb(ctx)
	newsFilter := models.NewsFilter{
		Page: models.Page{
			Rows: 5,
		},
	}
	listNews,_ := newsFilter.GetList(db)
	common.Success(ctx,iris.Map{
		"news": listNews,
	})
}