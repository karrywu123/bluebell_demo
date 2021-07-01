package controller

import (
	"bluebell_demo/dao/mysql"
	"bluebell_demo/models"
	"bluebell_demo/pkg/snowflake"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// 评论

// CommentHandler 创建评论
func CommentHandler(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		fmt.Println(err)
		ResponseError(c, CodeInvalidParams)
		return
	}

	// 生成帖子ID
	commentID, err := snowflake.GetID()
	if err != nil {
		zap.L().Error("snowflake.GetID() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 获取作者ID，当前请求的UserID
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	comment.CommentID = commentID
	comment.AuthorID = userID

	// 创建帖子
	if err := mysql.CreateComment(&comment); err != nil {
		zap.L().Error("mysql.CreatePost(&post) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// CommentListHandler 评论列表
func CommentListHandler(c *gin.Context) {
	postId := c.Param("id")
	ids,err:= strconv.ParseInt(postId, 10, 64)
	posts, err := mysql.GetCommentListByIDs(ids)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}

	for i :=range posts{
		str :=strconv.FormatUint(posts[i].Author_id, 10)
		dta,err :=mysql.GetUserByID(str)
		if err!=nil {
			fmt.Println(err)
		}
		posts[i].Username=dta.UserName
	}
	ResponseSuccess(c, posts)
}

