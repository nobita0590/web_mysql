package course

import (
	"gopkg.in/kataras/iris.v8"
	"github.com/nobita0590/web_mysql/modules/common"
)

func Init(newsRoute iris.Party)  {
	newsRoute.Options("/list",common.Options)
	newsRoute.Options("/detail",common.Options)
	newsRoute.Options("/register",common.Options)
	newsRoute.Options("/",common.Options)
	newsRoute.Get("/list",listCourse)
	newsRoute.Get("/",aCourse)
	newsRoute.Get("/detail",publicCourse)
	newsRoute.Get("/register",listRegisCourse)
	newsRoute.Post("/",insertCourse)
	newsRoute.Put("/",updateCourse)
	newsRoute.Delete("/",deleteCourse)
}