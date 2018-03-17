package question

import (
	"gopkg.in/kataras/iris.v8"
	"github.com/nobita0590/web_mysql/modules/common"
)

func Init(newsRoute iris.Party)  {
	// questions
	newsRoute.Options("/list",common.Options)
	newsRoute.Options("/",common.Options)
	newsRoute.Get("/list",getListQuestion)
	newsRoute.Get("/",getAQuestion)
	newsRoute.Post("/",insertQuestion)
	newsRoute.Put("/",updateQuestion)
	newsRoute.Delete("/",deleteQuestion)
	// tests
	newsRoute.Options("/tests/list",common.Options)
	newsRoute.Options("/tests/",common.Options)
	newsRoute.Options("/tests/exams",common.Options)
	newsRoute.Options("/tests/topten",common.Options)

	newsRoute.Get("/tests/list",getListTests)
	newsRoute.Get("/tests/",getATests)
	newsRoute.Post("/tests/",insertTests)
	newsRoute.Put("/tests/",updateTests)
	newsRoute.Delete("/tests/",deleteTests)
	newsRoute.Get("/tests/exams",getTestsExams)
	newsRoute.Get("/tests/topten",getTopTen)
	// tests frame
	newsRoute.Options("/tests/frame/list",common.Options)
	newsRoute.Options("/tests/frame",common.Options)
	newsRoute.Options("/tests/frame/generate",common.Options)

	newsRoute.Get("/tests/frame/list",getListTestsFrame)
	newsRoute.Get("/tests/frame",getATestsFrame)
	newsRoute.Post("/tests/frame",insertTestsFrame)
	newsRoute.Put("/tests/frame",updateTestsFrame)
	newsRoute.Delete("/tests/frame",deleteTestsFrame)
	newsRoute.Post("/tests/frame/generate",generateExamFrame)
}
