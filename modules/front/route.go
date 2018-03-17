package front

import (
	"gopkg.in/kataras/iris.v8"
	"github.com/nobita0590/web_mysql/modules/common"
	"github.com/nobita0590/web_mysql/modules/user"
	"github.com/nobita0590/web_mysql/modules/course"
)

func Init(frontRoute iris.Party)  {
	{
		frontRoute.Options("/home",common.Options)
		frontRoute.Options("/user",common.Options)
		frontRoute.Options("/news",common.Options)
		frontRoute.Options("/course/regis",common.Options)
		frontRoute.Get("/news",newsPage)
		frontRoute.Get("/home",homePage)
		frontRoute.Put("/user",common.NeedLoginMiddleWare,checkForEditUser,user.UpdateUser)
		frontRoute.Post("/course/regis",common.NeedLoginMiddleWare,course.RegisCourse)
	}
	/* exam route */
	{
		examRoute := frontRoute.Party("/exams")
		examRoute.Options("/",common.Options)
		examRoute.Post("/",common.NeedLoginMiddleWare,examInsert)
	}

	/* fags route */
	{
		fagsRoute := frontRoute.Party("/fags",common.NeedLoginMiddleWare)
		fagsRoute.Options("/",common.Options)
		fagsRoute.Options("/list",common.Options)
		fagsRoute.Get("/list",listFags)
		fagsRoute.Get("/",aFags)
		fagsRoute.Post("/",insertFags)
		fagsRoute.Put("/",updateFags)
		fagsRoute.Delete("/",deleteFags)

		/* comment route */
		routeComment := fagsRoute.Party("/comments")
		routeComment.Options("/",common.Options)
		routeComment.Options("/list",common.Options)
		routeComment.Options("/vote",common.Options)
		routeComment.Options("/trust",common.Options)

		routeComment.Get("/list",listComments)
		routeComment.Get("/",aComment)
		routeComment.Post("/",insertComment)
		routeComment.Put("/",updateComment)
		routeComment.Delete("/",deleteComment)
		routeComment.Post("/vote",voteComment)
		routeComment.Post("/trust",trustComment)
	}
}