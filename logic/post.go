package logic

import (
	"studyWeb/Bluebell/dao/mysql"
	"studyWeb/Bluebell/models"
	"studyWeb/Bluebell/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(post *models.Post) (err error) {
	//生成ID
	post.ID = int64(snowflake.GenID())
	//保存到数据库
	//返回参数
	return mysql.CreatePost(post)
}

func GetPostDetailByID(pid int64) (data *models.ApiPostDetail, err error) {
	post, err := mysql.GetPostByID(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostByID() failed", zap.Error(err))
		return
	}

	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID() failed", zap.Error(err))
		return
	}
	communityDetailByID, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID() failed", zap.Error(err))
		return nil, err
	}
	data = &models.ApiPostDetail{
		AuthorName: user.Username,
		//VoteNum:         0,
		Post:            post,
		CommunityDetail: communityDetailByID,
	}

	return
}
