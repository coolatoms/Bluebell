package controller

import (
	"errors"

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
