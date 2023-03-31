package redis_ser

import (
	"context"
	"encoding/json"
	"gbv2/config/es"
	"gbv2/config/log"
	"gbv2/models"
	"github.com/olivere/elastic/v7"
	"time"
)

func LikeSync() {
	for {
		// 获取文章
		result, err := es.ES.
			Search(models.ArticleModel{}.Index()).
			Query(elastic.NewMatchAllQuery()).
			Size(10000).
			Do(context.Background())

		if err != nil {
			log.Errorw(err.Error(), "err", err)
			return
		}

		diggInfo := GetDiggInfo()
		lookInfo := GetLookInfo()
		for _, hit := range result.Hits.Hits {
			var article models.ArticleModel
			// 反序列化为结构体
			err = json.Unmarshal(hit.Source, &article)

			digg := diggInfo[hit.Id]
			look := lookInfo[hit.Id]

			newDigg := article.DiggCount + digg
			newLook := article.LookCount + look
			// 如果点赞数不变
			if article.DiggCount == newDigg && article.LookCount == newLook {
				log.Infow(article.Title, "点赞数无变化与浏览量都不变", nil)
				continue
			}
			// 变换更新es
			_, err := es.ES.
				Update().
				Index(models.ArticleModel{}.Index()).
				Id(hit.Id).
				Doc(map[string]int{
					"digg_count": newDigg,
					"look_count": newLook,
				}).
				Do(context.Background())
			if err != nil {
				log.Errorw(err.Error(), "err", err)
				continue
			}
			log.Infow(article.Title, "点赞浏览数据同步成功, 点赞数", newDigg, "浏览数", newLook)
		}
		DiggClear()
		// 每天同步一次
		time.Sleep(24 * time.Hour)
	}

}
