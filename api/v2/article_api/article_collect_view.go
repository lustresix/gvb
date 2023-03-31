package article_api

import (
	"context"
	"encoding/json"
	"fmt"
	"gbv2/config/es"
	"gbv2/config/log"
	"gbv2/models"
	"gbv2/models/res"
	"gbv2/service/common"
	"gbv2/utils/jwt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type CollResponse struct {
	models.ArticleModel
	CreatedAt string `json:"created_at"`
}

func (ArticleApi) ArticleCollListView(c *gin.Context) {

	var cr models.PageInfo

	c.ShouldBindQuery(&cr)

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	var articleIDList []interface{}

	list, count, err := common.CommonList(models.UserCollectModel{UserID: claims.UserID}, common.Option{
		PageInfo: cr,
	})

	var collMap = map[string]string{}

	for _, model := range list {
		articleIDList = append(articleIDList, model.ArticleID)
		collMap[model.ArticleID] = model.CreatedAt.Format("2006-01-02 15:04:05")
	}

	boolSearch := elastic.NewTermsQuery("_id", articleIDList...)

	var collList = make([]CollResponse, 0)

	// 传id列表，查es
	result, err := es.ES.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	fmt.Println(result.Hits.TotalHits.Value, articleIDList)

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			log.Errorw(err.Error(), "err", err)
			continue
		}
		article.ID = hit.Id
		collList = append(collList, CollResponse{
			ArticleModel: article,
			CreatedAt:    collMap[hit.Id],
		})
	}
	res.OKWitList(collList, count, c)
}
