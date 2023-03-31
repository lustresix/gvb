package article_api

import (
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/res"
	"gbv2/service/es_ser"
	"gbv2/utils/jwt"
	"github.com/gin-gonic/gin"
)

// ArticleCollCreateView 用户收藏文章，或取消收藏
func (ArticleApi) ArticleCollCreateView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	model, err := es_ser.CommDetail(cr.ID)
	if err != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}

	var coll models.UserCollectModel
	err = mysql.DB.Take(&coll, "user_id = ? and article_id = ?", claims.UserID, cr.ID).Error
	var num = -1
	if err != nil {
		// 没有找到 收藏文章
		mysql.DB.Create(&models.UserCollectModel{
			UserID:    claims.UserID,
			ArticleID: cr.ID,
		})
		// 给文章的收藏数 +1
		num = 1
	}
	// 取消收藏
	// 文章数 -1
	mysql.DB.Delete(&coll)

	// 更新文章收藏数
	err = es_ser.ArticleUpdate(cr.ID, map[string]any{
		"collects_count": model.CollectsCount + num,
	})

	if num == 1 {
		res.OKWithMsg("收藏文章成功", c)
	} else {
		res.OKWithMsg("取消收藏成功", c)
	}
}
