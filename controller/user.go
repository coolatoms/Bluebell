package controller

import (
	"net/http"
	"studyWeb/Bluebell/logic"
	"studyWeb/Bluebell/models"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SingUpHandler(ctx *gin.Context) {

	//	1,获取参数参数校验
	var p models.ParamSignUp
	//判断类型有限
	err := ctx.ShouldBindJSON(&p)
	if err != nil {
		zap.L().Error("SignUp with invalid param ", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}

	//	2，业务处理
	logic.SingUp(&p)
	//	3，返回相应
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
