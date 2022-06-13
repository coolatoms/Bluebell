package controller

import (
	"strconv"
	"studyWeb/Bluebell/logic"
	"studyWeb/Bluebell/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CreatePostHandler 创建帖子
func CreatePostHandler(ctx *gin.Context) {
	//	1,获取参数及校验
	p := new(models.Post)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Debug(" ctx.ShouldBindJSON failed", zap.Any("err", err))
		zap.L().Error("ctx.ShouldBindJSON(p) failed")
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	//从ctx中获取
	userID, err := GetCurrentUser(ctx)
	if err != nil {
		ResponseError(ctx, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	//	2,创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreattPost() failed", zap.Error(err))
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	//	3,返回响应
	ResponseSuccess(ctx, "success")
}

// GetPostDetailHandler 获取帖子详情
func GetPostDetailHandler(ctx *gin.Context) {
	//获取参数，从url中获取
	pidStr := ctx.Param("id")

	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(ctx, CodeInvalidParam)
	}
	//根据ID查数据库
	data, err := logic.GetPostDetailByID(pid)
	if err != nil {
		zap.L().Error("logic.GetPostDetailByID(pid) failed ", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}

	//返回参数
	ResponseSuccess(ctx, data)
}

// GetPostListHandler 获取帖子列表
func GetPostListHandler(ctx *gin.Context) {
	//1，获取参数
	page, size := GetPageInfo(ctx)
	//2，数据库查询
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failes", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	//3，返回参数
	ResponseSuccess(ctx, data)
}

// GetPostListHandler2 获取帖子列表升级版
func GetPostListHandler2(ctx *gin.Context) {
	//1，获取参数
	//2，在redis中查询值
	p := &models.ParamPostList{
		CommunityID: 0,
		Page:        1,
		Size:        10,
		Order:       models.OrderTime,
	}
	err := ctx.ShouldBindQuery(p)
	if err != nil {
		zap.L().Error("GetPostListHandler2 with invalid param", zap.Error(err))
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	//业务处理
	data, err := logic.GetPostListNew(p)
	if err != nil {
		return
	}
	//3，根据ID去数据库中查ID

	//2，数据库查询

	if err != nil {
		zap.L().Error("logic.GetPostList() failes", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	//3，返回参数
	ResponseSuccess(ctx, data)
}
