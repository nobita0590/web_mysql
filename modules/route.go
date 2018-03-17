package modules

import (
	"gopkg.in/kataras/iris.v8"
	"github.com/nobita0590/web_mysql/modules/common"
	"github.com/nobita0590/web_mysql/modules/auth"
	"github.com/nobita0590/web_mysql/modules/user"
	"github.com/nobita0590/web_mysql/modules/news"
	"github.com/nobita0590/web_mysql/modules/document"
	"github.com/nobita0590/web_mysql/modules/question"

	"github.com/nobita0590/web_mysql/modules/u_config"
	"github.com/nobita0590/web_mysql/modules/front"
	"github.com/nobita0590/web_mysql/modules/course"
	"github.com/nobita0590/web_mysql/modules/upload"
	"strings"
)

var (
	fileCache CacheContent
)

func frondPath(ctx iris.Context) {
	file := fileCache.TakeContent("/public/dist/frontend.html")
	ctx.Write(file)
}

func BindRoute(app *iris.Application){
	fileCache.files = make(map[string][]byte)

	app.Options("/{id}", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"status": true,
		})
	})
	app.Get("/admin", func(ctx iris.Context) {
		file := fileCache.TakeContent("/public/dist/backend.html")
		ctx.Write(file)
	})
	app.Get("/admin/{path:path}", func(ctx iris.Context) {
		file := fileCache.TakeContent("/public/dist/backend.html")
		ctx.Write(file)
	})
	app.Get("/", frondPath)
	app.Get("/trang-chu", frondPath)
	app.Get("/khoa-hoc", frondPath)
	app.Get("/khoa-hoc/{path:path}", frondPath)
	app.Get("/tin-tuc", frondPath)
	app.Get("/tin-tuc/{path:path}", frondPath)
	app.Get("/nguoi-dung", frondPath)
	app.Get("/nguoi-dung/{path:path}", frondPath)
	app.Get("/tai-khoan", frondPath)
	app.Get("/tai-khoan/{path:path}", frondPath)
	app.Get("/thi-online", frondPath)
	app.Get("/thi-online/{path:path}", frondPath)
	app.Get("/hoi-dap", frondPath)
	app.Get("/hoi-dap/{path:path}", frondPath)
	app.Get("/tai-lieu", frondPath)
	app.Get("/tai-lieu/{path:path}", frondPath)

	app.Get("/{file:file}",func(ctx iris.Context){
		fileName := ctx.Params().Get("file")
		split := strings.Split(fileName,".")
		switch split[len(split) - 1] {
		case "css":
			ctx.Header("Content-Type","text/css; charset=utf-8")
		case "js":
			ctx.Header("Content-Type","application/javascript")
		default:
			ctx.Header("Content-Type","text/plain; charset=utf-8")
		}
		file := fileCache.TakeContent("/public/dist/"+ fileName)
		ctx.Write(file)
	})
	app.Get("/assets/{path:path}",func(ctx iris.Context){
		file := fileCache.TakeContent("/public/dist/"+ctx.Params().Get("path"))
		ctx.Write(file)
	})


	apiParty := app.Party("/api",common.InitMiddleWare)

	upload.Init(apiParty.Party("/upload"))
	auth.Use(apiParty.Party("/auth"))
	user.Init(apiParty.Party("/user",common.NeedLoginMiddleWare))
	news.Init(apiParty.Party("/news",common.NeedLoginMiddleWare))
	document.Init(apiParty.Party("/document",common.NeedLoginMiddleWare))
	question.Init(apiParty.Party("/question",common.NeedLoginMiddleWare))
	u_config.Init(apiParty.Party("/config",common.NeedLoginMiddleWare))
	course.Init(apiParty.Party("/course",common.NeedLoginMiddleWare))
	front.Init(apiParty.Party("/front"))
}
