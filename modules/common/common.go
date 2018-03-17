package common

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/kataras/iris.v8"

	"github.com/nobita0590/web_mysql/key"
	"github.com/nobita0590/web_mysql/models"
	"github.com/SermoDigital/jose/jwt"
	"encoding/json"
	"strings"
	"github.com/nobita0590/web_mysql/config"
)

const (
	DbKey = "DbKey"
	UserSession = "UserSession"
)

func Options(ctx iris.Context){
	ctx.JSON(iris.Map{"status":true})
}

func InitMiddleWare(ctx iris.Context){
	store := ctx.Values()
	// root:@Va123456@tcp(45.77.170.201)/study_db?charset=utf8&parseTime=True
	db, err := gorm.Open("mysql", config.MysqlConnect)
	defer db.Close()
	if(err != nil){
		ctx.JSON(iris.Map{
			"status": false,
			"error": "Cant not connect to database",
		})
		return
	}
	store.Set(DbKey,db)
	ctx.Header("Access-Control-Allow-Origin","*")//http://localhost:4201
	ctx.Header("Access-Control-Allow-Methods","GET, POST, PUT, OPTIONS, DELETE")
	ctx.Header("Access-Control-Allow-Headers","Authorization,X-XSRF-TOKEN,Content-Type")
	ctx.Header("Access-Control-Allow-Credentials", "true");
	ctx.Header("supports_credentials", "true");
	//store.Se
	ctx.Next()
}

func NeedLoginMiddleWare(ctx iris.Context)  {
	if ctx.Method() == "OPTIONS" {
		ctx.Next()
		return
	}
	var (
		err error
		userClaimByte,accesskey []byte
		claims jwt.Claims
		user  models.User
		userClaimString string
		ok bool
	)
	keyStr := ctx.GetHeader("Authorization")
	if (keyStr == "") {
		keyStr = ctx.FormValue("Authorization")
	}
	bearerSplit := strings.Split(keyStr," ")

	if !(len(bearerSplit) == 2 && bearerSplit[0] == "Bearer"){
		if access_token := ctx.FormValue("access_token");access_token != "" {
			accesskey = []byte(access_token)
		}else{
			goto Error
		}
	}else{
		accesskey = []byte(bearerSplit[1])
	}

	claims,err = key.ValidateKey(accesskey)
	userClaimString,ok =  claims.Get("user").(string)
	userClaimByte = []byte(userClaimString)

	if !ok {
		goto Error
	}
	if err != nil{
		goto Error
	}
	err = json.Unmarshal(userClaimByte,&user)
	if err != nil{
		goto Error
	}else{
		//Success(ctx,user)
		store := ctx.Values()
		store.Set(UserSession,user)
		ctx.Next()
		return
	}
Error:
	if ctx.Method() != "GET" {
		ctx.StatusCode(iris.StatusUnauthorized)
		errStr := ""
		if err != nil {
			errStr = err.Error()
		}
		ctx.JSON(iris.Map{
			"status" : false,
			"error" : "Key is not valid!",
			"detail" : errStr,
		})
	}else{
		store := ctx.Values()
		store.Set(UserSession,models.User{})
		ctx.Next()
	}
}

func GetDb(ctx iris.Context) (db *gorm.DB,err error) {
	store := ctx.Values()
	var ok bool
	if db,ok = store.Get(DbKey).(*gorm.DB);ok{
		return
	}else{
		err = errors.New("Can not to get Db")
		return
	}
}

func GetUser(ctx iris.Context) (user models.User,err error) {
	store := ctx.Values()
	var ok bool
	if user,ok = store.Get(UserSession).(models.User);ok{
		return
	}else{
		err = errors.New("Can not to get User")
		return
	}
}