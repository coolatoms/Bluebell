//router 注册路由
package router

import (
	"net/http"
	"studyWeb/Bluebell/controller"
	"studyWeb/Bluebell/logger"
	"studyWeb/Bluebell/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //发布模式
	}
	//r := gin.Default()
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	//注册
	v1 := r.Group("/api/v1")

	v1.POST("/signup", controller.SingUpHandler)
	//登录
	v1.POST("/login", controller.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/posts", controller.GetPostListHandler)
		//根据时间或分数获取帖子列表
		v1.GET("/post2", controller.GetPostListHandler2)

		v1.POST("/vote", controller.PostVoteHandler)
	}
	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
