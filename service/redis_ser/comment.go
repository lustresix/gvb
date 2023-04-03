package redis_ser

import (
	"gbv2/config/redis"
	"strconv"
)

const commentPrefix = "comment"

// Comment 点赞某一篇文章
func Comment(id string) error {
	num, _ := redis.RDB.HGet(commentPrefix, id).Int()
	num++
	err := redis.RDB.HSet(commentPrefix, id, num).Err()
	return err
}

// GetComment 获取某一篇文章下的点赞数
func GetComment(id string) int {
	num, _ := redis.RDB.HGet(commentPrefix, id).Int()
	return num
}

// GetCommentInfo 取出点赞数据
func GetCommentInfo() map[string]int {
	var DiggInfo = map[string]int{}
	maps := redis.RDB.HGetAll(commentPrefix).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		DiggInfo[id] = num
	}
	return DiggInfo
}

func CommentClear() {
	redis.RDB.Del(commentPrefix)
}
