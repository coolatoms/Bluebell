package redis

import (
	"strconv"
	"studyWeb/Bluebell/models"
	"time"

	"github.com/go-redis/redis"
)

func getIDsFormKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + size - 1
	//ZRevRange 查询
	return client.ZRevRange(key, start, end).Result()
}

func GetPostIDsInORder(list *models.ParamPostList) ([]string, error) {
	//根据用户选择确定需要的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if list.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	//确定查询的起始点
	return getIDsFormKey(key, list.Page, list.Size)
}

// GetPostVoteData 根据ids查询每篇帖子的数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	//data = make([]int64, 0, len(ids))
	//for _, id := range ids {
	//	key := getRedisKey(KeyPostVotedZSetPF + id)
	//	//统计帖子赞成票的数量
	//	v1 := client.ZCount(key, "1", "1").Val()
	//	data = append(data, v1)
	//}
	pipline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPF + id)
		pipline.ZCount(key, "1", "1")
	}
	exec, err := pipline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(exec))
	for _, cmder := range exec {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return

}

// GetCommunityPostListByIDs 根据社区查询帖子参数
func GetCommunityPostListByIDs(list *models.ParamPostList) ([]string, error) {
	//根据用户选择确定需要的redis key
	orderkey := getRedisKey(KeyPostTimeZSet)
	if list.Order == models.OrderScore {
		orderkey = getRedisKey(KeyPostScoreZSet)
	}
	key := orderkey + strconv.Itoa(int(list.CommunityID))

	ckey := KeyCommunitySetPF + strconv.Itoa(int(list.CommunityID))

	if client.Exists(orderkey).Val() < 1 {
		pipline := client.Pipeline()
		pipline.ZInterStore(key, redis.ZStore{
			Weights:   nil,
			Aggregate: "MAX",
		}, getRedisKey(ckey), orderkey)
		pipline.Expire(key, 60*time.Second)
		_, err := pipline.Exec()
		if err != nil {
			return nil, err
		}
	}
	return getIDsFormKey(key, list.Page, list.Size)
}
