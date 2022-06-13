package controller

import (
	"errors"
	"fmt"
	"studyWeb/Bluebell/dao/mysql"
	"studyWeb/Bluebell/logic"
	"studyWeb/Bluebell/models"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// SingUpHandler 注册逻辑
func SingUpHandler(ctx *gin.Context) {

	//	1,获取参数参数校验
	var p models.ParamSignUp
	//判断类型有限
	err := ctx.ShouldBindJSON(&p)
	if err != nil {
		zap.L().Error("SignUp with invalid param ", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(ctx, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//	2，业务处理
	err = logic.SingUp(&p)
	if err != nil {
		zap.L().Error("Logic.singUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(ctx, CodeUserExist)
		}
		ResponseError(ctx, CodeServerBusy)
		return
	}
	//	3，返回相应
	ResponseSuccess(ctx, nil)
}

// LoginHandler 登录逻辑
func LoginHandler(ctx *gin.Context) {
	//	1，获取请求参数
	p := new(models.ParamLogin) //new返回的是指针
	err := ctx.ShouldBindJSON(p)
	if err != nil {
		zap.L().Error("Login with invalid param ", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(ctx, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	//	2，业务逻辑
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.login failed", zap.String("username", p.Username), zap.Error(err))
		ResponseError(ctx, CodeInvalidPassword)
		return
	}
	//	3，返回相应
	ResponseSuccess(ctx, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID),
		"user_name": user.Username,
		"token":     user.Token,
	})
}
