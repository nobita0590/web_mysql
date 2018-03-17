package document

import (
	"gopkg.in/kataras/iris.v8"
	"github.com/nobita0590/web_mysql/modules/common"
)

func Init(newsRoute iris.Party)  {
	newsRoute.Options("/list",common.Options)
	newsRoute.Options("/front",common.Options)
	newsRoute.Options("/",common.Options)

	newsRoute.Get("/list",getListDocument)
	newsRoute.Get("/front",getFrontDocument)
	newsRoute.Get("/",getADocument)
	newsRoute.Get("/download/{id:int}",downloadDocument)
	newsRoute.Get("/preview/{id:int}",previewDocument)
	newsRoute.Post("/",insertDocument)
	newsRoute.Put("/",updateDocument)
	newsRoute.Delete("/",deleteDocument)
}
