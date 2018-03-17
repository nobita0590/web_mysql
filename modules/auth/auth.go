package auth

import (
	"gopkg.in/kataras/iris.v8"

	"github.com/nobita0590/web_mysql/modules/common"
	"github.com/nobita0590/web_mysql/models"

	"golang.org/x/crypto/bcrypt"

	"github.com/SermoDigital/jose/jws"

	"errors"

	"github.com/nobita0590/web_mysql/key"
	"encoding/json"
	"github.com/SermoDigital/jose/jwt"
)

func Use(authRoute iris.Party)  {
	authRoute.Options("/register",common.Options)
	authRoute.Post("/register",registerUser)
	authRoute.Options("/login",common.Options)
	authRoute.Post("/login",login)
	authRoute.Options("/login/facebook",common.Options)
	authRoute.Post("/login/facebook",loginFacebook)
	authRoute.Options("/login/google",common.Options)
	authRoute.Post("/login/google",loginGoogle)
	authRoute.Get("/valid",valid)

}

func registerUser(ctx iris.Context)  {
	db,_ := common.GetDb(ctx)
	user := models.User{}
	if err,validErrs := common.ReadStruct(ctx,&user,true);err == nil{
		uniqueUser := models.User{}
		if err := db.Where("email = ?",user.Email).Where("social_id = ?","").
			First(&uniqueUser).Error;err == nil{
			// found user
			detailValid := make(map[string]string)
			if uniqueUser.Email == user.Email {
				detailValid["Email"] = "The Email has used by other"
			}
			err = errors.New("The field is unique")
			common.BadRequest(ctx,err,detailValid)
		}else{
			//common.Success(ctx,uniqueUser)
			newPass,_ := bcrypt.GenerateFromPassword([]byte(user.Password),11)
			user.Password = string(newPass)
			err := db.Create(&user).Error
			if(err == nil){
				accessToken,err := generateAccessToken(user)

				common.Success(ctx,iris.Map{
					"user": user,
					"err": err,
					"access_token": accessToken,
				})
			}else{
				common.InternalServer(ctx,err)
			}
		}
	}else{
		common.BadRequest(ctx,err,validErrs)
	}
}

func login(ctx iris.Context){
	email := ctx.FormValue("Email")
	password := ctx.FormValue("Password")
	db,_ := common.GetDb(ctx)
	user := models.User{}
	if err := db.Where("email = ?",email).First(&user).Error;err == nil{
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));err == nil{
			accessToken,err := generateAccessToken(user)

			common.Success(ctx,iris.Map{
				"user": user,
				"err": err,
				"access_token": accessToken,
			})
		}else{
			// wrong password
			common.BadRequest(ctx,err,iris.Map{
				"Password": "Mật khẩu không đúng",
			})
		}
	}else{
		//not found
		common.BadRequest(ctx,err,iris.Map{
			"Email": "Tài khoản không hợp lệ",
		})
	}
}

func generateAccessToken(user models.User) (string,error) {
	claimUser,_ := json.Marshal(user)
	claims := jws.Claims{
		"user": string(claimUser),
	}
	accessToken,err := key.GenerateKey(claims)
	return string(accessToken),err
}

func valid(ctx iris.Context)  {
	var (
		err error
		userClaimByte []byte
		claims jwt.Claims
		user  models.User
	)
	accesskey := []byte(ctx.FormValue("access_token"))
	claims,err = key.ValidateKey(accesskey)
	userClaimString,ok :=  claims.Get("user").(string)
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
		common.Success(ctx,user)
		return
	}

Error:
	common.BadRequest(ctx,err,"Key is not valid!")
}


