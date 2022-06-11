//router 注册路由
package router

import (
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
	r.POST("/signup", controller.SingUpHandler)
	r.POST("/login", controller.LoginHandler)
	r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		//	如果是登陆的用户，判断请求头中是否包含，有效的JWT，认证放到中间件中middlewares.JWTAuthMiddleware()
		c.Request.Header.Get("Authorization")
	})

	return r
}
