package front

import (
	"gopkg.in/kataras/iris.v8"
	"github.com/nobita0590/web_mysql/models"
	"github.com/nobita0590/web_mysql/modules/common"
	"strconv"
	"errors"
)

func homePage(ctx iris.Context) {
	db,_ := common.GetDb(ctx)
	/*  news  */
	newsFilter := models.NewsFilter{
		Page: models.Page{
			Rows: 5,
		},
	}
	listNews,_ := newsFilter.GetList(db)
	/* course */
	courseFilter := models.CourseFilter{
		Page: models.Page{
			Rows: 10,
		},
	}
	courses,_ := courseFilter.GetList(db.New())
	if len(courses) > 4 && len(courses) < 8 {
		courses = courses[:4]
	}
	/* docs */
	docFilter := models.DocumentFilter{
		Page: models.Page{
			Rows: 12,
		},
	}
	docs,_ := docFilter.GetList(db.New())
	/* fags */
	fagFilter := models.FagsFilter{}
	newFags,_ := fagFilter.GetList(db.New(),0)
	hotFags,_ := fagFilter.GetList(db.New(),0)
	/* exam */
	selectSourceFilter := models.SelectSourceFilter{
		GroupsId:[]uint{2},
		Page: models.Page{Rows:4},
	}
	subjects,_ := selectSourceFilter.GetList(db.New())
	banner,_ := models.BannerFilter{ID:1}.GetA(db.New())

	common.Success(ctx,iris.Map{
		"news": listNews,
		"courses": courses,
		"docs": docs,
		"newFags": newFags,
		"hotFags": hotFags,
		"subjects": subjects,
		"banners": banner.InfoView,
	})
}

func checkForEditUser(ctx iris.Context) {
	if user,err := common.GetUser(ctx);err == nil {
		if id, err := strconv.ParseUint(ctx.FormValue("ID"), 10, 64); err == nil {
			if user.ID == uint(id) {
				ctx.Next()
				return
			}
		}
	}
	common.Unauthorized(ctx,errors.New("Bạn không được phép"))
}