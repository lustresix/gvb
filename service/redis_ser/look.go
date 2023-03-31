package redis_ser

import (
	"gbv2/config/redis"
	"strconv"
)

const lookPrefix = "look"

// Look 浏览某一篇文章
func Look(id string) error {
	num, _ := redis.RDB.HGet(lookPrefix, id).Int()
	num++
	err := redis.RDB.HSet(lookPrefix, id, num).Err()
	return err
}

// GetLook 获取某一篇文章下的浏览数
func GetLook(id string) int {
	num, _ := redis.RDB.HGet(lookPrefix, id).Int()
	return num
}

// GetLookInfo 取出浏览量数据
func GetLookInfo() map[string]int {
	var DiggInfo = map[string]int{}
	maps := redis.RDB.HGetAll(lookPrefix).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		DiggInfo[id] = num
	}
	return DiggInfo
}

func LookClear() {
	redis.RDB.Del(lookPrefix)
}
