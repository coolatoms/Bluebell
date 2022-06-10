//router 注册路由
package router

import (
	"studyWeb/Bluebell/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	//r:=gin.New()
	//r.Use(logger.GinLogger(),logger.GinRecovery())
	r.POST("/signup", controller.SingUpHandler)
	return r
}
