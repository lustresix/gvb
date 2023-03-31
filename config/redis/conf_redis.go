package redis

import (
	"context"
	"gbv2/config/log"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"time"
)

var RDB *redis.Client

func RedisInit() {
	RDB = ConnectRedis()
}

func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"), // no password set
		DB:       viper.GetInt("redis.db"),          // use default DB
		PoolSize: viper.GetInt("redis.pool_size"),   // 连接池大小
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		log.Errorw("redis连接失败", "err", err)
		return nil
	}
	log.Infow("redis连接成功")
	// 同步点赞数
	return rdb
}
