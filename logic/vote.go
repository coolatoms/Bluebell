package logic

import (
	"strconv"
	"studyWeb/Bluebell/dao/redis"
	"studyWeb/Bluebell/models"

	"go.uber.org/zap"
)

//投票功能
//1，用户投票数据

func PostVote(userid int64, p *models.ParamVoteData) error {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userid),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userid)), p.PostID, float64(p.Direction))
}
