package controller

import (
	"bluebell_demo/dao/redis"
	"encoding/json"
	"fmt"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type VoteData struct {
	PostID    string  `json:"post_id"`
	Direction float64 `json:"direction,string"`
}

func (v *VoteData) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		PostID    string  `json:"post_id"`
		Direction float64 `json:"direction,string"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.PostID) == 0 {
		err = errors.New("缺少必填字段post_id")
	} else if required.Direction == 0 {
		err = errors.New("缺少必填字段direction")
	} else {
		v.PostID = required.PostID
		v.Direction = required.Direction
	}
	return
}

func VoteHandler(c *gin.Context) {
	// 给哪个文章投什么票
	var vote VoteData
	if err := c.ShouldBindJSON(&vote); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNotLogin)
		return
	}
	if err := redis.PostVote(vote.PostID, fmt.Sprint(userID), vote.Direction); err != nil {
		zap.L().Error("投票报错", zap.String("投票ID", vote.PostID), zap.Error(err))
		ResponseErrorWithMsg(c,CodeServerBusy,fmt.Sprintf("%v",err))
		return
	}
	ResponseSuccess(c, "vote is ok")
}
