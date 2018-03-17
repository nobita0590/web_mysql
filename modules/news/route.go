package news

import (
	"gopkg.in/kataras/iris.v8"
	"github.com/nobita0590/web_mysql/modules/common"
)

func Init(newsRoute iris.Party)  {
	newsRoute.Options("/list",common.Options)
	newsRoute.Options("/root",common.Options)
	newsRoute.Options("/detail",common.Options)
	newsRoute.Options("/",common.Options)
	newsRoute.Get("/root",getRoot)
	newsRoute.Get("/list",listNews)
	newsRoute.Get("/",aNews)
	newsRoute.Get("/detail",publicNews)
	newsRoute.Post("/",insertNews)
	newsRoute.Put("/",updateNews)
	newsRoute.Delete("/",deleteNews)
}
