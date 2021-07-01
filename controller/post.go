package controller

import (
	"bluebell_demo/dao/mysql"
	"bluebell_demo/dao/redis"
	"bluebell_demo/logic"
	"bluebell_demo/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// PostHandler 创建帖子
func CreatePostHandler(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	// 参数校验

	// 获取作者ID，当前请求的UserID
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	post.AuthorId = userID

	err = logic.CreatePost(&post)
	if err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// PostListHandler 帖子列表
func PostListHandler(c *gin.Context) {
	order, _ := c.GetQuery("order")
	pageStr, ok := c.GetQuery("page")
	fmt.Println(pageStr)
	if !ok {
		pageStr = "1"
	}
	pageNum, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		pageNum = 1
	}
	if order=="time"{
		posts,_:=mysql.GetPostList1(pageNum)
		ResponseSuccess(c, posts)
	}
	posts := redis.GetPost(order, pageNum)
	ResponseSuccess(c, posts)
}

func PostList2Handler(c *gin.Context) {
	data, err := logic.GetPostList2()
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)

}

// PostDetailHandler 帖子详情
func PostDetailHandler(c *gin.Context) {
	postId := c.Param("id")

	post, err := logic.GetPost(postId)
	if err != nil {
		zap.L().Error("logic.GetPost(postID) failed", zap.String("postId", postId), zap.Error(err))
	}

	ResponseSuccess(c, post)
}

