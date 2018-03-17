package upload

import (
	"gopkg.in/kataras/iris.v8"
	"github.com/nobita0590/web_mysql/modules/common"
	"io"
	"os"
	"strings"
	"github.com/nobita0590/web_mysql/config"
)
var(
	FileUploadPath = config.FilePath + "/public/uploads"
	PrivateUploadPath = config.FilePath + "/private/"
)

type Filter struct {
	Command		string		`form:"command"`
	Lang		string		`form:"lang"`
	Type		string		`form:"type"`
	CurrentFolder		string		`form:"currentFolder"`
	Hash		string		`form:"hash"`
	Token		string		`form:"token"`
	ResponseType		string		`form:"responseType"`
}
func Init(confRoute iris.Party)  {
	confRoute.Options("/user",common.Options)
	confRoute.Post("/user", common.NeedLoginMiddleWare ,
		iris.LimitRequestBodySize(10<<20), uploadUserImage)

	confRoute.Options("/",common.Options)
	confRoute.Post("/", iris.LimitRequestBodySize(10<<20), uploadImage)

	confRoute.Options("/banner",common.Options)
	confRoute.Post("/banner", iris.LimitRequestBodySize(10<<20), uploadBanner)

	confRoute.Options("/private",common.Options)
	confRoute.Post("/private", iris.LimitRequestBodySize(10<<40), uploadPrivateFile)

	confRoute.Options("/ckfinder",common.Options)
	confRoute.Post("/ckfinder", iris.LimitRequestBodySize(10<<20), ckfinder)

	confRoute.Options("/ckeditor",common.Options)
	confRoute.Post("/ckeditor",ckeditor)
}

func generateFileName(oldFileName string) (newFileName string) {
	p := common.NewProvider()
	split := strings.Split(oldFileName,".")
	newFileName = p.Generate(15) + "." + split[len(split) - 1]
	return
}

func ckeditor(ctx iris.Context){
	file, info, err := ctx.FormFile("upload")

	if err != nil {
		common.InternalServer(ctx,err)
		return
	}

	defer file.Close()

	fileName := generateFileName(info.Filename)

	// Create a file with the same name
	// assuming that you have a folder named 'uploads'
	out, err := os.OpenFile(FileUploadPath+ "/"+fileName,
		os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		common.InternalServer(ctx,err)
		return
	}else{
		CKEditorFuncNum := ctx.FormValue("CKEditorFuncNum")
		ctx.HTML(`
<script type="text/javascript">
    window.parent.CKEDITOR.tools.callFunction(` + CKEditorFuncNum +`,"` + `http://` + config.Host + "/public/uploads/" + fileName+`" , "");
</script>
	`)
	//"http://localhost:8080/public/uploads/1509212804_rar.png"
	}
	defer out.Close()
	io.Copy(out, file)
	// CKEditor := ctx.FormValue("CKEditor")

	// langCode := ctx.FormValue("langCode")

}

func uploadUserImage(ctx iris.Context) {
	user,err := common.GetUser(ctx);
	if err != nil {
		common.Unauthorized(ctx,err)
		return
	}
	file, info, err := ctx.FormFile("image")

	if err != nil {
		common.InternalServer(ctx,err)
		return
	}

	defer file.Close()
	fileName := generateFileName(info.Filename)

	// Create a file with the same name
	// assuming that you have a folder named 'uploads'
	out, err := os.OpenFile(FileUploadPath+ "/user/"+fileName,
		os.O_WRONLY|os.O_CREATE, 0666)

	db,_ := common.GetDb(ctx)
	db.Model(&user).Updates(map[string]interface{}{
		"avatar_url": "/public/uploads/user/" + fileName,
	})

	if err != nil {
		common.InternalServer(ctx,err)
		return
	}else{
		common.Success(ctx,iris.Map{
			"FileName" : fileName,
			"FilePath" : "/public/uploads/user/" + fileName,
		})
	}
	defer out.Close()
	io.Copy(out, file)
}

func uploadImage(ctx iris.Context) {
	file, info, err := ctx.FormFile("image")

	if err != nil {
		common.InternalServer(ctx,err)
		return
	}

	defer file.Close()
	fileName := generateFileName(info.Filename)

	// Create a file with the same name
	// assuming that you have a folder named 'uploads'
	out, err := os.OpenFile(FileUploadPath+ "/"+fileName,
		os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		common.InternalServer(ctx,err)
		return
	}else{
		common.Success(ctx,iris.Map{
			"FileName" : fileName,
			"FilePath" : "/public/uploads/" + fileName,
		})
	}
	defer out.Close()
	io.Copy(out, file)
}

func uploadBanner(ctx iris.Context) {
	file, info, err := ctx.FormFile("image")

	if err != nil {
		common.InternalServer(ctx,err)
		return
	}

	defer file.Close()
	fileName := generateFileName(info.Filename)

	// Create a file with the same name
	// assuming that you have a folder named 'uploads'
	out, err := os.OpenFile(FileUploadPath+ "/banner/"+fileName,
		os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		common.InternalServer(ctx,err)
		return
	}else{
		common.Success(ctx,iris.Map{
			"FileName" : fileName,
			"FilePath" : "/public/uploads/banner/" + fileName,
		})
	}
	defer out.Close()
	io.Copy(out, file)
}

func uploadPrivateFile(ctx iris.Context) {
	file, info, err := ctx.FormFile("private_file")

	if err != nil {
		common.InternalServer(ctx,err)
		return
	}

	defer file.Close()
	fileName := generateFileName(info.Filename)

	// Create a file with the same name
	// assuming that you have a folder named 'uploads'
	out, err := os.OpenFile(PrivateUploadPath + fileName,
		os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		common.InternalServer(ctx,err)
		return
	}else{
		common.Success(ctx,iris.Map{
			"FileName" : fileName,
			"FilePath" : "/private/" + fileName,
		})
	}
	defer out.Close()
	io.Copy(out, file)
}

func ckfinder(ctx iris.Context)  {
	f := Filter{}
	if err,_ := common.ReadStruct(ctx,&f,false);err == nil {
		switch f.Command {
		case "Init":
			ctx.JSON(iris.Map{
				"c": "",
				"enabled": true,
				"images": iris.Map{
					"max": "1600x1200",
					"sizes": iris.Map{
						"large": "800x600",
						"medium": "600x480",
						"small": "480x320",
					},
				},
				"resourceTypes": []iris.Map{
					{
						"acl": 1023,
						"allowedExtensions": "7z,aiff,asf,avi,bmp,csv,doc,docx,fla,flv,gif,gz,gzip,jpeg,jpg,mid,mov,mp3,mp4,mpc,mpeg,mpg,ods,odt,pdf,png,ppt,pptx,pxd,qt,ram,rar,rm,rmi,rmvb,rtf,sdc,sitd,swf,sxc,sxw,tar,tgz,tif,tiff,txt,vsd,wav,wma,wmv,xls,xlsx,zip",
						"deniedExtensions": "",
						"hasChildren": true,
						"hash": "0d8612c92f0bf497",
						"maxSize": 2097152,
						"name": "Files",
						"url": "/ckfinder/userfiles/files/",
					},{
						"acl": 1023,
						"allowedExtensions": "bmp,gif,jpeg,jpg,png",
						"deniedExtensions": "",
						"hasChildren": true,
						"hash": "6380d1d3f6f653a3",
						"maxSize": 2097152,
						"name": "Images",
						"url": "/ckfinder/userfiles/images/",
					},
				},
			})
			return
		case "GetFolders":
		case "GetFiles":
		case "FileUpload":


		}
	}else {
		
	}
}
