package redis_ser

import (
	"encoding/json"
	"fmt"
	"gbv2/config/redis"
	"time"
)

const newsIndex = "news_index"

type NewData struct {
	Index    string `json:"index"`
	Title    string `json:"title"`
	HotValue string `json:"hotValue"`
	Link     string `json:"link"`
}

// SetNews 设置某一个数据，重复执行，重复累加
func SetNews(key string, newData []NewData) error {
	byteData, _ := json.Marshal(newData)
	err := redis.RDB.Set(fmt.Sprintf("%s_%s", newsIndex, key), byteData, 10*time.Second).Err()
	return err
}

func GetNews(key string) (newData []NewData, err error) {
	res := redis.RDB.Get(fmt.Sprintf("%s_%s", newsIndex, key)).Val()
	err = json.Unmarshal([]byte(res), &newData)
	return
}
