package logic

import (
	"studyWeb/Bluebell/dao/mysql"
	"studyWeb/Bluebell/dao/redis"
	"studyWeb/Bluebell/models"
	"studyWeb/Bluebell/pkg/snowflake"

	"go.uber.org/zap"
)

// CreatePost 创建帖子
func CreatePost(post *models.Post) (err error) {
	//生成ID
	post.ID = int64(snowflake.GenID())
	//保存到数据库
	err = redis.CreatePost(post.ID, post.CommunityID)
	if err != nil {
		return err
	}
	//返回参数
	return mysql.CreatePost(post)
}

// GetPostDetailByID 通过ID获取帖子
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
		return
	}
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: communityDetailByID,
	}

	return
}

// GetPostList 获取帖子列表
func GetPostList(page int64, size int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		zap.L().Error("mysql.GetPostList() failed", zap.Error(err))
		return
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID() failed", zap.Error(err))
			continue
		}
		communityDetailByID, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID() failed", zap.Error(err))
			continue
		}

		postPetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: communityDetailByID,
		}
		data = append(data, postPetail)
	}
	return
}

// GetPostList2 根据给定的id列表查询相应的数据
func GetPostList2(list *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	ids, err := redis.GetPostIDsInORder(list)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn(" redis.GetPostIDsInORder return zero data")
		return
	}
	//	根据id去数据库中查询相应的帖子详细信息
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))
	//查好帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return nil, err
	}
	for idx, post := range posts {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID() failed", zap.Error(err))
			continue
		}
		communityDetailByID, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID() failed", zap.Error(err))
			continue
		}

		postPetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: communityDetailByID,
		}
		data = append(data, postPetail)
	}

	return

}

func GetCommunityPostList(list *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	ids, err := redis.GetCommunityPostListByIDs(list)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn(" redis.GetPostIDsInORder return zero data")
		return
	}
	//	根据id去数据库中查询相应的帖子详细信息
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))
	//查好帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return nil, err
	}
	for idx, post := range posts {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID() failed", zap.Error(err))
			continue
		}
		communityDetailByID, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID() failed", zap.Error(err))
			continue
		}

		postPetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: communityDetailByID,
		}
		data = append(data, postPetail)
	}
	return

}

// GetPostListNew 根据给定ID或communityID查询接口
func GetPostListNew(list *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	if list.CommunityID == 0 {
		//查所有
		data, err = GetPostList2(list)
	} else {
		//根据社区ID查询
		data, err = GetCommunityPostList(list)
	}
	if err != nil {
		zap.L().Error("logic.GetPostListNew failed", zap.Error(err))
		return
	}
	return
}
