package article_api

import (
	"gbv2/config/log"
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/ctype"
	"gbv2/models/res"
	"gbv2/service/es_ser"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"time"
)

type ArticleUpdateRequest struct {
	Title    string   `json:"title"`     // 文章标题
	Abstract string   `json:"abstract"`  // 文章简介
	Content  string   `json:"content"`   // 文章内容
	Category string   `json:"category"`  // 文章分类
	Source   string   `json:"source"`    // 文章来源
	Link     string   `json:"link"`      // 原文链接
	BannerID uint     `json:"banner_id"` // 文章封面id
	Tags     []string `json:"tags"`      // 文章标签
	ID       string   `json:"id"`
}

func (ArticleApi) ArticleUpdateView(c *gin.Context) {
	var cr ArticleUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.Errorw(err.Error(), "err", err)
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}

	var coverModel models.ImageModel
	var coverUrl string
	if cr.BannerID != 0 {
		err = mysql.DB.Model(models.ImageModel{}).Where("id = ?", cr.BannerID).Scan(&coverModel).Error
		if err != nil {
			res.FailWithMsg("图片不存在", c)
			return
		}
		coverUrl = coverModel.Path + "/" + coverModel.Name
	}

	article := models.ArticleModel{
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		Title:     cr.Title,
		Keyword:   cr.Title,
		Abstract:  cr.Abstract,
		Content:   cr.Content,
		Category:  cr.Category,
		Source:    cr.Source,
		Link:      cr.Link,
		BannerID:  cr.BannerID,
		BannerUrl: coverUrl,
		Tags:      cr.Tags,
	}

	maps := structs.Map(&article)
	var DataMap = map[string]any{}
	// 去掉空值
	for key, v := range maps {
		switch val := v.(type) {
		case string:
			if val == "" {
				continue
			}
		case uint:
			if val == 0 {
				continue
			}
		case int:
			if val == 0 {
				continue
			}
		case ctype.Array:
			if len(val) == 0 {
				continue
			}
		case []string:
			if len(val) == 0 {
				continue
			}
		}
		DataMap[key] = v
	}

	err = article.GetDataByID(cr.ID)
	if err != nil {
		log.Errorw(err.Error(), "err", err)
		res.FailWithMsg("文章不存在", c)
		return
	}

	err = es_ser.ArticleUpdate(cr.ID, DataMap)
	if err != nil {
		log.Errorw(err.Error(), "err", err)
		res.FailWithMsg("更新失败", c)
		return
	}

	model, _ := es_ser.CommDetail(cr.ID)
	if article.Content != model.Content || article.Title != model.Title {
		es_ser.DeleteFullTextByArticleID(cr.ID)
		es_ser.AsyncArticleByText(cr.ID, article.Title, model.Content)
	}
	res.OKWithMsg("更新成功", c)
}
