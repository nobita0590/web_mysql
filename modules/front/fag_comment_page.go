package front

import (
	"gopkg.in/kataras/iris.v8"
	"github.com/nobita0590/web_mysql/models"
	"github.com/nobita0590/web_mysql/modules/common"
	"github.com/kataras/iris/core/errors"
	"encoding/json"
)
func listComments(ctx iris.Context) {
	filter := models.FagCommentsFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		user,_ := common.GetUser(ctx)
		if comments,err := filter.GetList(db,user.ID);err == nil{
			common.Success(ctx,iris.Map{
				"models": comments,
				"p_info": filter.GetPageInfo(),
			})
			//common.Success(ctx,listNews)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}

func aComment(ctx iris.Context) {
	filter := models.FagCommentsFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if comment,err := filter.GetA(db);err == nil{
			common.Success(ctx,comment)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}

func insertComment(ctx iris.Context) {
	db,_ := common.GetDb(ctx)
	var comment models.FagComments
	if err,validErrs := common.ReadJSONStruct(ctx,&comment,true);err == nil{
		user,_ := common.GetUser(ctx)
		comment.UserId = user.ID
		voteInfo,_ := json.Marshal(models.FagCommentVoteInfo{})
		comment.VoteInfo = string(voteInfo)
		if err := db.Create(&comment).Error;err == nil{
			db.Exec("UPDATE fags SET comments_number = comments_number + 1 WHERE id = ?",comment.FagId)
			common.Success(ctx,comment)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,validErrs)
	}
}

func updateComment(ctx iris.Context) {
	db,_ := common.GetDb(ctx)
	var comment models.FagComments
	if err,validErrs := common.ReadJSONStruct(ctx,&comment,true);err == nil{
		db.LogMode(true)
		if err := db.Model(&comment).Updates(map[string]interface{}{
			"content": comment.Content,
		}).Error;err == nil{
			common.Success(ctx,comment.ID)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,validErrs)
	}
}

func voteComment(ctx iris.Context) {
	var filter models.FagCommentsFilter
	var comment models.FagComments
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil {
		db,_ := common.GetDb(ctx)
		if err := db.First(&comment,filter.ID).Error;err == nil {
			user,_ := common.GetUser(ctx)
			if user.ID == comment.UserId {
				common.Unauthorized(ctx,errors.New("Khong duoc phep"))
			}else{
				voteInfo := models.FagCommentVoteInfo{}
				json.Unmarshal([]byte(comment.VoteInfo),&voteInfo)
				if voteValid := comment.UpdateVote(voteInfo,user.ID,filter.IsTrusted);voteValid {
					db.LogMode(true)
					if err := db.Model(&comment).Update(map[string]interface{}{
						"upvote": comment.Upvote,"downvote": comment.Downvote,"vote_info": comment.VoteInfo,
					}).Error;err == nil{
						common.Success(ctx,comment)
					}else{
						common.InternalServer(ctx,err)
					}
				}else{
					common.Unauthorized(ctx,errors.New("Ban da thuc hien thao tac nay roi"))
				}
			}
		}else{
			common.NotFound(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}

func trustComment(ctx iris.Context) {
	user,_ := common.GetUser(ctx)
	var filter models.FagCommentsFilter
	var fag models.Fags
	var comment models.FagComments
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil {
		db,_ := common.GetDb(ctx)
		err := db.First(&fag,filter.FagId).Error
		if err == nil {
			if !user.IsAdmin && user.ID != fag.UserId {
				common.Unauthorized(ctx,errors.New("Ban khong duoc phep truy cap"));
				return
			}
			if err = db.First(&comment,filter.ID).Error;err == nil{
				validUpdate := false
				if !fag.IsDone && filter.IsTrusted {
					fag.IsDone = true
					comment.IsTrusted = true
					validUpdate = true
				}
				if fag.IsDone && !filter.IsTrusted {
					fag.IsDone = false
					comment.IsTrusted = false
					validUpdate = true
				}
				if validUpdate {
					db.Model(&fag).Where("id = ?", fag.ID).Updates(map[string]interface{}{
						"is_done": fag.IsDone,
					})
					db.Model(&comment).Where("id = ?", comment.ID).Updates(map[string]interface{}{
						"is_trusted": comment.IsTrusted,
					})
					common.Success(ctx,filter.IsTrusted)
				}else{
					common.BadRequest(ctx,errors.New("Thao tac khong hop le"),nil)
				}
				return
			}
		}
		common.BadRequest(ctx,err,nil)
	}else{
		common.BadRequest(ctx,err,nil)
	}
}
func deleteComment(ctx iris.Context) {
	filter := models.FagCommentsFilter{}
	if err,_ := common.ReadStruct(ctx,&filter,false);err == nil{
		db,_ := common.GetDb(ctx)
		if err := filter.Delete(db);err == nil{
			common.Success(ctx,filter.ID)
		}else{
			common.InternalServer(ctx,err)
		}
	}else{
		common.BadRequest(ctx,err,nil)
	}
}
