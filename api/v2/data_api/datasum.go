package data_api

import (
	"context"
	"gbv2/config/es"
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/res"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type DataSumResponse struct {
	UserCount      int `json:"user_count"`
	ArticleCount   int `json:"article_count"`
	MessageCount   int `json:"message_count"`
	ChatGroupCount int `json:"chat_group_count"`
	NowLoginCount  int `json:"now_login_count"`
	NowSignCount   int `json:"now_sign_count"`
}

func (DataApi) DataSumView(c *gin.Context) {

	var userCount, articleCount, messageCount, ChatGroupCount int
	var nowLoginCount, nowSignCount int

	result, _ := es.ES.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Do(context.Background())
	articleCount = int(result.Hits.TotalHits.Value) //搜索到结果总条数
	mysql.DB.Model(models.UserModel{}).Select("count(id)").Scan(&userCount)
	mysql.DB.Model(models.MessageModel{}).Select("count(id)").Scan(&messageCount)
	mysql.DB.Model(models.ChatModel{IsGroup: true}).Select("count(id)").Scan(&ChatGroupCount)
	mysql.DB.Model(models.LoginDataModel{}).Where("to_days(created_at)=to_days(now())").
		Select("count(id)").Scan(&nowLoginCount)
	mysql.DB.Model(models.UserModel{}).Where("to_days(created_at)=to_days(now())").
		Select("count(id)").Scan(&nowSignCount)

	res.OKWithData(DataSumResponse{
		UserCount:      userCount,
		ArticleCount:   articleCount,
		MessageCount:   messageCount,
		ChatGroupCount: ChatGroupCount,
		NowLoginCount:  nowLoginCount,
		NowSignCount:   nowSignCount,
	}, c)
}
