package auth

import (
	"gopkg.in/kataras/iris.v8"
	fb "github.com/huandu/facebook"
	"github.com/nobita0590/web_mysql/modules/common"
	"errors"
	"fmt"
	"reflect"
	"github.com/nobita0590/web_mysql/models"
	"time"
)

var (
	FbAppId = "1192039274236789"
	FbAppSecret = "143b42deac26288a493690bcbe80bbf3"
	invalidToken = errors.New("Invalid access token")
)

type (
	FbResponse struct {
		Id 		string		`json:"id"`
		About	string		`json:"about"`
		FirstName	string	`json:"first_name"`
		LastName	string	`json:"last_name"`
		Education	[]FbEduInfo
		Cover	struct{
			Id 		string		`json:"id"`
			OffsetX		float32		`json:"offset_x"`
			OffsetY		float32		`json:"offset_y"`
			Source		string		`json:"source"`
		}		`json:"cover"`
		Birthday	string		`json:"birthday"`
		Email		string		`json:"email"`
		Gender		string		`json:"gender"`
		Picture		struct{
			Data		struct{
				IsSilhouette		bool		`json:"is_silhouette"`
				Url					string		`json:"url"`
			}		`json:"data"`
		}		`json:"picture"`
	}
	FbEduInfo	struct {
		School	struct{
			Id 		string		`json:"id"`
			Name	string		`json:"name"`
		}	`json:"school"`
		Type 	string		`json:"type"`
		Id 		string		`json:"id"`
	}
	LongTermFbTokenResponse struct {
		AccessToken 	string		`json:"access_token"`
		TokenType 		string		`json:"token_type"`
		ExpiresIn		uint		`json:"expires_in"`
	}
)

func loginFacebook(ctx iris.Context) {
	token := ctx.FormValue("access_token")
	id := ctx.FormValue("id")
	if res, err := fb.Get("/oauth/access_token", fb.Params{
		"client_id": FbAppId,
		"client_secret": FbAppSecret,
		"grant_type": "fb_exchange_token",
		"fb_exchange_token": token,
	});err == nil {
		longToken := LongTermFbTokenResponse{}
		if err := common.ConvertThrowJson(res,&longToken);err == nil {
			db,_ := common.GetDb(ctx)
			filter := models.UserFilter{
				SocialId: "fb_" + id,
			}
			if user,err := filter.GetA(db);err == nil {
				generateLoginSocial(ctx,user)
				return
			}else{
				userInfo, _ := fb.Get("/"+id, fb.Params{
					"fields": "id,about,first_name,last_name,education,cover,birthday,email,gender,picture",
					"access_token": longToken.AccessToken,
				})
				//fmt.Println(userInfo)
				if _,ok := userInfo["error"];ok {
					common.BadRequest(ctx,invalidToken,iris.Map{
						"1": userInfo,
						"tk": longToken,
						"u_id": "/"+id,
					})
				}else{
					fbResponse := FbResponse{}
					if err := common.ConvertThrowJson(userInfo,&fbResponse);err == nil{
						user := fbResponse.ParseToUser()
						oldUser := models.User{}
						if err = db.Where("social_id = ?",user.SocialId).First(&oldUser).Error;err == nil {
							/*if user.SocialId == oldUser.SocialId {
								generateLoginSocial(ctx,user)
							}else{
								common.BadRequest(ctx,errors.New("Email đã có người sử dụng"),nil)
							}*/
							generateLoginSocial(ctx,user)
						}else{
							err := db.Create(&user).Error
							if(err == nil){
								generateLoginSocial(ctx,user)
							}else{
								common.InternalServer(ctx,err)
							}
						}
					}else{
						common.BadRequest(ctx,invalidToken,2)
					}
				}
			}
		}else{
			common.BadRequest(ctx,invalidToken,3)
		}
	}else{
		common.BadRequest(ctx,invalidToken,4)
	}
}

func generateLoginSocial(ctx iris.Context, user models.User)  {
	accessToken,err := generateAccessToken(user)
	common.Success(ctx,iris.Map{
		"user": user,
		"err": err,
		"access_token": accessToken,
	})
}

func loginGoogle(ctx iris.Context) {
	/*token := ctx.FormValue("access_token")
	id := ctx.FormValue("id")
	if res, err := fb.Get("/oauth/access_token", fb.Params{
		"client_id": FbAppId,
		"client_secret": FbAppSecret,
		"grant_type": "fb_exchange_token",
		"fb_exchange_token": token,
	});err == nil {
		longToken := LongTermFbTokenResponse{}
		if err := common.ConvertThrowJson(res,&longToken);err == nil {
			db,_ := common.GetDb(ctx)
			filter := models.UserFilter{
				SocialId: "fb_" + id,
			}
			if user,err := filter.GetA(db);err == nil {
				accessToken,err := generateAccessToken(user)
				common.Success(ctx,iris.Map{
					"user": user,
					"err": err,
					"access_token": accessToken,
				})
			}else{
				userInfo, _ := fb.Get("/"+id, fb.Params{
					"fields": "id,about,first_name,last_name,education,cover,birthday,email,gender,picture",
					"access_token": longToken.AccessToken,
				})
				fmt.Println(userInfo)
				if _,ok := userInfo["error"];ok {
					common.BadRequest(ctx,invalidToken,iris.Map{
						"1": userInfo,
						"tk": longToken,
						"u_id": "/"+id,
					})
				}else{
					fbResponse := FbResponse{}
					if err := common.ConvertThrowJson(userInfo,&fbResponse);err == nil{
						user := fbResponse.ParseToUser()
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
					}else{
						common.BadRequest(ctx,invalidToken,2)
					}
				}
			}
		}else{
			common.BadRequest(ctx,invalidToken,3)
		}
	}else{
		common.BadRequest(ctx,invalidToken,4)
	}*/
}

func getLongToken(res fb.Result) (token string,err error) {
	if tokenInterface,ok := res["access_token"];ok {
		fmt.Println(reflect.TypeOf(tokenInterface))
		if reflect.TypeOf(tokenInterface).Kind() == reflect.String {
			token = tokenInterface.(string)
			return
		}
	}
	err = errors.New("unknown token")
	return
}

func (f FbResponse) ParseToUser() (user models.User) {
	user.FirstName = f.FirstName
	user.LastName = f.LastName
	user.SocialId = "fb_"+f.Id
	user.Email = f.Email
	user.AvatarUrl = f.Picture.Data.Url
	user.CorverUrl = f.Cover.Source
	for _,edu := range f.Education {
		if edu.Type == "High School" {
			user.School = edu.School.Name
		}
	}
	if f.Gender == "male" {
		user.Gender = 1
	}else{
		user.Gender = 0
	}
	if birthDay,err := time.Parse("02/01/2006",f.Birthday);err == nil{
		user.Birthday = birthDay
	}
	user.Description = f.About
	return
}