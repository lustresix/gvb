package article_api

import (
	"context"
	"encoding/json"
	"gbv2/config/es"
	"gbv2/config/log"
	"gbv2/models"
	"gbv2/models/res"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"time"
)

type CalendarResponse struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type BucketsType struct {
	Buckets []struct {
		KeyAsString string `json:"key_as_string"`
		Key         int64  `json:"key"`
		DocCount    int    `json:"doc_count"`
	} `json:"buckets"`
}

var DateCount = map[string]int{}

func (ArticleApi) ArticleCalendarView(c *gin.Context) {

	agg := elastic.NewDateHistogramAggregation().
		Field("created_at"). // 根据created_at字段值，对数据进行分组
		//  分组间隔：month代表每月、支持minute（每分钟）、hour（每小时）、day（每天）、week（每周）、year（每年)
		CalendarInterval("day")

	// 时间段搜索
	// 从今天开始，到去年的今天
	now := time.Now()
	aYearAgo := now.AddDate(-1, 0, 0)

	format := "2006-01-02 15:04:05"
	// lt 小于  gt 大于
	query := elastic.NewRangeQuery("created_at").
		Gte(aYearAgo.Format(format)).
		Lte(now.Format(format))

	result, err := es.ES.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("calendar", agg).
		Size(0).
		Do(context.Background())

	if err != nil {
		log.Errorw("err", "查询失败", err)
		res.FailWithMsg("查询失败", c)
		return
	}

	var data BucketsType
	_ = json.Unmarshal(result.Aggregations["calendar"], &data)

	var resList = make([]CalendarResponse, 0)
	for _, bucket := range data.Buckets {
		Time, _ := time.Parse(format, bucket.KeyAsString)
		DateCount[Time.Format("2006-01-02")] = bucket.DocCount
	}
	days := int(now.Sub(aYearAgo).Hours() / 24)
	for i := 0; i <= days; i++ {
		day := aYearAgo.AddDate(0, 0, i).Format("2006-01-02")

		count, _ := DateCount[day]
		resList = append(resList, CalendarResponse{
			Date:  day,
			Count: count,
		})
	}

	res.OKWithData(resList, c)

}
