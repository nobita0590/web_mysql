package user

import (
	"gopkg.in/kataras/iris.v8"
	"github.com/nobita0590/web_mysql/modules/common"
	"github.com/nobita0590/web_mysql/models"
	"golang.org/x/crypto/bcrypt"
	"errors"
)

func Init(userRoute iris.Party)  {
	userRoute.Options("/list",common.Options)
	userRoute.Options("/",common.Options)
	userRoute.Options("/change-password",common.Options)
	userRoute.Get("/list",listUser)
	userRoute.Get("/",aUser)
	userRoute.Post("/",insertUser)
	userRoute.Put("/",UpdateUser)
	userRoute.Delete("/",deleteUser)
	userRoute.Put("/change-password",ChangePassword)
}

func listUser(ctx iris.Context)  {
	filter := models.UserFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if users,err := filter.GetList(db);err == nil{
			common.Success(ctx,iris.Map{
				"user": users,
				"p_info": filter.GetPageInfo(),
			})
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}

func aUser(ctx iris.Context)  {
	filter := models.UserFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if user,err := filter.GetA(db);err == nil{
			common.Success(ctx,user)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}

func insertUser(ctx iris.Context)  {
	user := models.User{}
	if err,validErrs := common.ReadStruct(ctx,&user,true);err == nil{
		db,_ := common.GetDb(ctx)
		uniqueUser := models.User{}
		if err := db.Where("email = ?",user.Email).First(&uniqueUser).Error;err == nil{
			// found user
			detailValid := make(map[string]string)
			if uniqueUser.Email == user.Email {
				detailValid["Email"] = "The Email has used by other"
			}
			err = errors.New("The field is unique")
			common.BadRequest(ctx,err,detailValid)
		}else{
			newPass,_ := bcrypt.GenerateFromPassword([]byte(user.Password),11)
			user.Password = string(newPass)
			err := db.Create(&user).Error
			if(err == nil){
				common.Success(ctx,user)
			}else{
				common.InternalServer(ctx,err)
			}
		}
	}else{
		common.BadRequest(ctx,err,validErrs)
	}
}
func UpdateUser(ctx iris.Context)  {
	db,_ := common.GetDb(ctx)
	var newUser models.User
	if err,validErrs := common.ReadStruct(ctx,&newUser,true,"Password");err == nil{
		newUser.Password = ""
		newUser.Email = ""
		if err := db.Model(&newUser).Updates(map[string]interface{}{
			"first_name": newUser.FirstName,
			"last_name": newUser.LastName,
			"is_admin": newUser.IsAdmin,
			//"avatar_url": newUser.AvatarUrl,
			"phone": newUser.Phone,
			"birthday": newUser.Birthday,
			"class": newUser.Class,
			"school": newUser.School,
			"province_id": newUser.ProvinceId,
		}).Error;err == nil{
			common.Success(ctx,newUser.ID)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,validErrs)
	}
}

func ChangePassword(ctx iris.Context)  {
	db,_ := common.GetDb(ctx)
	user,_ := common.GetUser(ctx)
	db.First(&user, user.ID)
	oldPass := ctx.PostValue("old_password")
	password := ctx.PostValue("password")
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(oldPass));err == nil{
		newPass,_ := bcrypt.GenerateFromPassword([]byte(password),11)
		if err := db.Model(&user).Updates(map[string]interface{}{
			"password": string(newPass),
		}).Error;err == nil{
			common.Success(ctx,user.ID)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,iris.Map{
			"password": "Mật khẩu không đúng",
		})
	}
}
func deleteUser(ctx iris.Context)  {
	filter := models.UserFilter{}
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
func test(ctx iris.Context)  {
	ctx.JSON(ctx.Request().UserAgent())
}
