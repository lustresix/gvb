package article_api

import (
	"context"
	"encoding/json"
	"fmt"
	"gbv2/config/es"
	"gbv2/config/log"
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/res"
	"gbv2/service/es_ser"
	"gbv2/utils/jwt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

func (ArticleApi) ArticleCollBatchRemoveView(c *gin.Context) {
	var cr models.ESIDListRequest

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	var collects []models.UserCollectModel
	var articleIDList []string
	mysql.DB.Find(&collects, "user_id = ? and article_id in ?", claims.UserID, cr.IDList).
		Select("article_id").
		Scan(&articleIDList)
	if len(articleIDList) == 0 {
		res.FailWithMsg("请求非法", c)
		return
	}
	var idList []interface{}
	for _, s := range articleIDList {
		idList = append(idList, s)
	}
	// 更新文章数
	boolSearch := elastic.NewTermsQuery("_id", idList...)
	result, err := es.ES.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			log.Errorw(err.Error(), "err", err)
			continue
		}
		count := article.CollectsCount - 1
		err = es_ser.ArticleUpdate(hit.Id, map[string]any{
			"collects_count": count,
		})
		if err != nil {
			log.Errorw(err.Error(), "err", err)
			continue
		}
	}
	mysql.DB.Delete(&collects)
	res.OKWithMsg(fmt.Sprintf("成功取消收藏 %d 篇文章", len(articleIDList)), c)

}
