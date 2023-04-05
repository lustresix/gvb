package article_api

import (
	"context"
	"encoding/json"
	"gbv2/config/es"
	"gbv2/models"
	"gbv2/models/res"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

func (ArticleApi) FullTextSearchView(c *gin.Context) {
	var cr models.PageInfo
	_ = c.ShouldBindQuery(&cr)

	boolQuery := elastic.NewBoolQuery()

	if cr.Key != "" {
		boolQuery.Must(elastic.NewMultiMatchQuery(cr.Key, "title", "body"))
	}

	result, err := es.ES.
		Search(models.FullTextModel{}.Index()).
		Query(boolQuery).
		Highlight(elastic.NewHighlight().Field("body")).
		Size(100).
		Do(context.Background())
	if err != nil {
		return
	}
	count := result.Hits.TotalHits.Value //搜索到结果总条数
	fullTextList := make([]models.FullTextModel, 0)
	for _, hit := range result.Hits.Hits {
		var model models.FullTextModel
		json.Unmarshal(hit.Source, &model)

		body, ok := hit.Highlight["body"]
		if ok {
			model.Body = body[0]
		}

		fullTextList = append(fullTextList, model)
	}

	res.OKWitList(fullTextList, count, c)
}
