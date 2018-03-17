package document

import (
	"gopkg.in/kataras/iris.v8"
	"github.com/nobita0590/web_mysql/models"
	"github.com/nobita0590/web_mysql/modules/common"
	"strings"
	"io/ioutil"
	"github.com/nobita0590/web_mysql/config"
	"fmt"
)

func getListDocument(ctx iris.Context) {
	filter := models.DocumentFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if listNews,err := filter.GetList(db);err == nil{
			common.Success(ctx,iris.Map{
				"models": listNews,
				"p_info": filter.GetPageInfo(),
			})
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}

func getADocument(ctx iris.Context) {
	filter := models.DocumentFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if document,err := filter.GetA(db);err == nil{
			common.Success(ctx,document)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}

func getFrontDocument(ctx iris.Context) {
	filter := models.DocumentFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if document,err := filter.GetAForFront(db);err == nil{
			common.Success(ctx,document)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}

func insertDocument(ctx iris.Context) {
	db,_ := common.GetDb(ctx)
	var document models.Document
	if err,validErrs := common.ReadStruct(ctx,&document,true);err == nil{
		user,_ := common.GetUser(ctx)
		document.CreatorId = user.ID
		if err := db.Create(&document).Error;err == nil{
			common.Success(ctx,document.ID)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,validErrs)
	}
}

func downloadDocument(ctx iris.Context)  {
	id, _ := ctx.Params().GetInt("id")
	filter := models.DocumentFilter{
		Id: uint(id),
	}
	db,_ := common.GetDb(ctx)
	if document,err := filter.GetA(db);err == nil{
		file := config.FilePath + document.PathStore
		fileSpit := strings.Split(file,"/")
		db.Exec("UPDATE documents SET download_number = download_number + 1 WHERE id = ?",document.ID)
		ctx.SendFile(file, fileSpit[len(fileSpit) - 1])
		// common.Success(ctx,document)
	}else{
		common.InternalServer(ctx,err)
	}
}

func previewDocument(ctx iris.Context)  {
	id, _ := ctx.Params().GetInt("id")
	filter := models.DocumentFilter{
		Id: uint(id),
	}
	db,_ := common.GetDb(ctx)
	if document,err := filter.GetA(db);err == nil{
		file,err := ioutil.ReadFile(config.FilePath + document.PathStore)
		fmt.Println(err)
		ctx.Write(file)
	}else{
		common.InternalServer(ctx,err)
	}
}

func updateDocument(ctx iris.Context) {
	db,_ := common.GetDb(ctx)
	var document models.Document
	if err,validErrs := common.ReadStruct(ctx,&document,true);err == nil{
		if err := db.Model(&document).Updates(document).Error;err == nil{
			common.Success(ctx,document.ID)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,validErrs)
	}
}
func deleteDocument(ctx iris.Context) {
	filter := models.DocumentFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if err := filter.Delete(db);err == nil{
			common.Success(ctx,filter.Id)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}
