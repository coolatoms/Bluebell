package controller

import (
	"studyWeb/Bluebell/logic"
	"studyWeb/Bluebell/models"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func PostVoteHandler(ctx *gin.Context) {

	p := new(models.ParamVoteData)
	err := ctx.ShouldBindJSON(p)
	if err != nil {
		zap.L().Error("ctx.ShouldBindJSON", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(ctx, CodeInvalidParam, errData)
		return
	}
	userid, err := GetCurrentUser(ctx)
	if err != nil {
		ResponseError(ctx, CodeNeedLogin)
	}
	err = logic.PostVote(userid, p)
	if err != nil {
		zap.L().Error(" logic.PostVote failed", zap.Error(err))
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	ResponseSuccess(ctx, 123)
}
