package u_config

import (
	"gopkg.in/kataras/iris.v8"
	"github.com/nobita0590/web_mysql/modules/common"
)

func Init(confRoute iris.Party)  {
	confRoute.Options("/select",common.Options)
	confRoute.Options("/select/list",common.Options)
	confRoute.Options("/select/order",common.Options)
	confRoute.Options("/select/group",common.Options)

	confRoute.Get("/select",getASelect)
	confRoute.Post("/select",insertSelect)
	confRoute.Put("/select",updateSelect)
	confRoute.Delete("/select",deleteSelect)
	confRoute.Get("/select/list",getListSelect)
	confRoute.Get("/select/group",getListAndGroupSelect)
	confRoute.Put("/select/order",updateSelectOrder)

	confRoute.Options("/sub-category",common.Options)
	confRoute.Options("/sub-category/list",common.Options)
	confRoute.Options("/sub-category/order",common.Options)
	confRoute.Options("/sub-category/group",common.Options)

	confRoute.Get("/sub-category",getASubCategory)
	confRoute.Post("/sub-category",insertSubCategory)
	confRoute.Put("/sub-category",updateSubCategory)
	confRoute.Delete("/sub-category",deleteSubCategory)
	confRoute.Get("/sub-category/list",getListSubCategory)
	confRoute.Get("/sub-category/group",getListAndGroupSubCategory)
	confRoute.Put("/sub-category/order",updateSubCategoryOrder)

	confRoute.Options("/banner",common.Options)
	confRoute.Options("/banner/list",common.Options)

	confRoute.Get("/banner",getABanner)
	confRoute.Put("/banner",updateBanner)
	confRoute.Get("/banner/list",getListBanner)
}
