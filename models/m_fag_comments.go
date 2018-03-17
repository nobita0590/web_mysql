package models

import (
	"github.com/jinzhu/gorm"
	"encoding/json"
	"time"
)

type (
	FagComments struct {
		ID 			uint
		FagId		uint
		UserId		uint
		Content		string
		IsTrusted	bool
		Upvote		int
		Downvote	int
		CurrentUser		int			`gorm:"-"`
		VoteInfo	string			`json:"-"`
		UserName	string			`gorm:"-"`
		AvatarUrl 	string			`gorm:"-"`
		CreatedAt	time.Time
	}
	FagCommentsFilter struct {
		Page
		ID			uint
		FagId		uint
		IsTrusted	bool
	}
	FagCommentVoteInfo struct {
		Up 		[]uint
		Down	[]uint
	}
)

func (f *FagCommentsFilter) GetList(db *gorm.DB,userId uint) (comments []FagComments,err error) {
	// offset := f.getOffset()
	rows,err := db.Table("fag_comments").
		Select(`fag_comments.*,
				CONCAT(users.last_name,' ',users.first_name) user_name,users.avatar_url`).
		Joins("left join users on users.id = fag_comments.user_id").
		Where("fag_comments.fag_id = ?", f.FagId).
		/*Limit(f.Rows).Offset(offset).*/
		Order("fag_comments.is_trusted desc,fag_comments.upvote desc").Rows()
	if err != nil {
		return
	}
	comment := FagComments{}
	defer rows.Close()
	for rows.Next() {
		comment = FagComments{}
		err = db.ScanRows(rows, &comment)
		if err != nil {
			return
		}
		voteInfo := FagCommentVoteInfo{}
		json.Unmarshal([]byte(comment.VoteInfo),&voteInfo)
		for _,uId := range voteInfo.Up {
			if uId == userId {
				comment.CurrentUser = 1
				break
			}
		}
		if comment.CurrentUser == 0 {
			for _,uId := range voteInfo.Down {
				if uId == userId {
					comment.CurrentUser = -1
					break
				}
			}
		}
		comments = append(comments,comment)
	}
	//Find(&comments).Error
	if f.Count && err == nil {
		err = db.Model(&Fags{}).Count(&f.Total).Error
	}
	return
}

func(f FagCommentsFilter) GetA(db *gorm.DB) (fag FagComments,err error) {
	err = db.Table("fag_comments").
		Select(`fag_comments.*,
				class.value class_name,
				subject.value subject_name,
				CONCAT(users.last_name,' ',users.first_name) user_name,users.avatar_url`).
		Joins("left join users on users.id = fags.user_id").
		Where("fags.id = ?",f.ID).
		Limit(1).First(&fag).Error
	return
}

func(f FagCommentsFilter) Delete(db *gorm.DB) error {
	return db.Where("id = ?",f.ID).Delete(&FagComments{}).Error
}

func (c *FagComments) UpdateVote(vote FagCommentVoteInfo,userId uint,isUp bool) bool {
	if isUp {
		for _,id := range vote.Up {
			if userId == id {
				return false
			}
		}
		c.Upvote += 1
		c.CurrentUser = 1
		vote.Up = append(vote.Up,userId)
		for index,id := range vote.Down {
			if userId == id {
				vote.Down = append(vote.Down[0:index],vote.Down[index + 1:]...)
				c.Downvote -= 1
				break
			}
		}
	}else{
		for _,id := range vote.Down {
			if userId == id {
				return false
			}
		}
		c.Downvote += 1
		c.CurrentUser = -1
		vote.Down = append(vote.Down,userId)
		for index,id := range vote.Up {
			if userId == id {
				vote.Up = append(vote.Up[0:index],vote.Up[index + 1:]...)
				c.Upvote -= 1
				break
			}
		}
	}
	voteByte,_ := json.Marshal(vote)
	c.VoteInfo = string(voteByte)
	return true
}