package redis_ser

import (
	"gbv2/config/redis"
	"strconv"
)

type CountDB struct {
	Index string // 索引
}

// Set 设置某一个数据，重复执行，重复累加
func (c CountDB) Set(id string) error {
	num, _ := redis.RDB.HGet(c.Index, id).Int()
	num++
	err := redis.RDB.HSet(c.Index, id, num).Err()
	return err
}

// Get 获取某个的数据
func (c CountDB) Get(id string) int {
	num, _ := redis.RDB.HGet(c.Index, id).Int()
	return num
}

// GetInfo 取出数据
func (c CountDB) GetInfo() map[string]int {
	var DiggInfo = map[string]int{}
	maps := redis.RDB.HGetAll(c.Index).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		DiggInfo[id] = num
	}
	return DiggInfo
}

func (c CountDB) Clear() {
	redis.RDB.Del(c.Index)
}
