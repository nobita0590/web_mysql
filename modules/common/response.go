package common

import "gopkg.in/kataras/iris.v8"

func Success(context iris.Context,data interface{})  {
	context.JSON(iris.Map{
		"status" : true,
		"data" : data,
	})
}
func Unauthorized(context iris.Context,err error)  {
	//context.StatusCode(iris.Status)
	context.StatusCode(iris.StatusUnauthorized)
	context.JSON(iris.Map{
		"status" : false,
		"error" : err.Error(),
	})
}
func BadRequest(context iris.Context,err error,detail interface{}) {
	context.StatusCode(iris.StatusBadRequest)
	//iris.StatusUnauthorized
	context.JSON(iris.Map{
		"status" : false,
		"error" : err.Error(),
		"detail" : detail,
	})
}
func UnprocessableEntity(context iris.Context,err error,detail interface{}) {
	context.StatusCode(iris.StatusUnprocessableEntity)
	//iris.StatusUnauthorized
	context.JSON(iris.Map{
		"status" : false,
		"error" : err.Error(),
		"detail" : detail,
	})
}
func InternalServer(context iris.Context,err error)  {
	context.StatusCode(iris.StatusInternalServerError)
	context.JSON(iris.Map{
		"status" : false,
		"error" : err.Error(),
	})
}

func NotFound(context iris.Context,err error)  {
	context.StatusCode(iris.StatusNotFound)
	context.JSON(iris.Map{
		"status" : false,
		"error" : err.Error(),
	})
}