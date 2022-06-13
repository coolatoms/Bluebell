package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

const ConTextUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("用户未登录")

// GetCurrentUser 获取当前用户ID
func GetCurrentUser(ctx *gin.Context) (userID int64, err error) {
	i, ok := ctx.Get(ConTextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = i.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

// GetPageInfo 获取页面参数
func GetPageInfo(ctx *gin.Context) (int64, int64) {
	offsetstr := ctx.Query("page")
	limitstr := ctx.Query("size")
	offset, err := strconv.ParseInt(offsetstr, 10, 64)
	if err != nil {
		offset = 1
	}
	limit, err := strconv.ParseInt(limitstr, 10, 64)
	if err != nil {
		limit = 10
	}
	return offset, limit
}
