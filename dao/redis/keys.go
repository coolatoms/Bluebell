package redis

const (
	Prefix             = "bluebell:"   // 项目key前缀
	KeyPostTimeZSet    = "post:time"   // zset;贴子及发帖时间
	KeyPostScoreZSet   = "post:score"  // zset;贴子及投票的分数
	KeyPostVotedZSetPF = "post:voted:" // zset;记录用户及投票类型;参数是post id

	KeyCommunitySetPF = "community:" //set;保存每个分支下的帖子ID
)

// 给redis key加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}
