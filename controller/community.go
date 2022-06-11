package controller

import (
	"strconv"
	"studyWeb/Bluebell/logic"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CommunityHandler 社区列表查询
func CommunityHandler(ctx *gin.Context) {
	//	1，查询到所有的社区（community_id,community_name)以列表的形式返回数据
	list, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	//
	ResponseSuccess(ctx, list)
}

// CommunityDetailHandler 社区分类详情
func CommunityDetailHandler(ctx *gin.Context) {
	//获取社区ID
	communityId := ctx.Param("id")
	id, err := strconv.ParseInt(communityId, 10, 64)
	if err != nil {
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	list, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	//
	ResponseSuccess(ctx, list)
}
