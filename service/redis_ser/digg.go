package redis_ser

import (
	"gbv2/config/redis"
	"strconv"
)

const diggPrefix = "digg"

// Digg 点赞某一篇文章
func Digg(id string) error {
	num, _ := redis.RDB.HGet(diggPrefix, id).Int()
	num++
	err := redis.RDB.HSet(diggPrefix, id, num).Err()
	return err
}

// GetDigg 获取某一篇文章下的点赞数
func GetDigg(id string) int {
	num, _ := redis.RDB.HGet(diggPrefix, id).Int()
	return num
}

// GetDiggInfo 取出点赞数据
func GetDiggInfo() map[string]int {
	var DiggInfo = map[string]int{}
	maps := redis.RDB.HGetAll(diggPrefix).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		DiggInfo[id] = num
	}
	return DiggInfo
}

func DiggClear() {
	redis.RDB.Del(diggPrefix)
}
