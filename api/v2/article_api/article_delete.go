package article_api

import (
	"context"
	"fmt"
	"gbv2/config/es"
	"gbv2/config/log"
	"gbv2/models"
	"gbv2/models/res"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type IDListRequest struct {
	IDList []string `json:"id_list"`
}

func (ArticleApi) ArticleRemoveView(c *gin.Context) {
	var cr IDListRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.Errorw(err.Error(), "err", err)
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}

	bulkService := es.ES.Bulk().Index(models.ArticleModel{}.Index()).Refresh("true")
	for _, id := range cr.IDList {
		req := elastic.NewBulkDeleteRequest().Id(id)
		bulkService.Add(req)
	}
	result, err := bulkService.Do(context.Background())
	if err != nil {
		log.Errorw(err.Error(), "err", err)
		res.FailWithMsg("删除失败", c)
		return
	}
	res.OKWithMsg(fmt.Sprintf("成功删除 %d 篇文章", len(result.Succeeded())), c)
	return

}
