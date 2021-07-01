package routers

import (
	"bluebell_demo/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//r.Use(logger.GinLogger(),logger.GinRecovery(true))
	//注册业务路由
	r.POST("/testsigup",controller.Test_SignUpHandler)
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	v1 := r.Group("/api/v1")
	v1.POST("/login", controller.LoginHandler)
	v1.POST("/signup", controller.SignUpHandler)
	v1.GET("/refresh_token", controller.RefreshTokenHandler)

//
	v1.Use(controller.JWTAuthMiddleware())
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.PostDetailHandler)
		v1.GET("/post", controller.PostListHandler)

		v1.GET("/post2", controller.PostList2Handler)

		v1.POST("/vote", controller.VoteHandler)

		v1.POST("/comment", controller.CommentHandler)
		v1.GET("/comment/:id", controller.CommentListHandler)

		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})

	}
//


	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}