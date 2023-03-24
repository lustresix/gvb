package article_api

import (
	"fmt"
	"gbv2/config/log"
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/ctype"
	"gbv2/models/res"
	"gbv2/utils/jwt"
	"github.com/gin-gonic/gin"
	"time"
)

type ArticleRequest struct {
	Title    string      `json:"title" binding:"required" msg:"文章标题必填"`   // 文章标题
	Abstract string      `json:"abstract"`                                      // 文章简介
	Content  string      `json:"content" binding:"required" msg:"文章内容必填"` // 文章内容
	Category string      `json:"category"`                                      // 文章分类
	Source   string      `json:"source"`                                        // 文章来源
	Link     string      `json:"link"`                                          // 原文链接
	BannerID uint        `json:"banner_id"`                                     // 文章封面id
	Tags     ctype.Array `json:"tags"`                                          // 文章标签
}

func (ArticleApi) ArticleCreateView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	userID := claims.UserID
	userNickName := claims.NickName
	fmt.Println("token获取成功")
	var cr ArticleRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.Errorw("err", "err", err)
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}

	// 截取前 30 个汉字
	if cr.Abstract == "" {
		abs := []rune(cr.Content)
		if len(abs) > 30 {
			cr.Abstract = string(abs[:30])
		} else {
			cr.Abstract = string(abs[:])
		}

	}

	var coverModel models.ImageModel
	err = mysql.DB.Model(models.ImageModel{}).Where("id = ?", cr.BannerID).Scan(&coverModel).Error
	if err != nil {
		res.FailWithMsg("图片不存在", c)
		return
	}
	coverUrl := coverModel.Path + "/" + coverModel.Name

	// 用户头像
	var avatar string
	err = mysql.DB.Model(models.UserModel{}).Where("avatar = ?", userID).Scan(&avatar).Error
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	article := models.ArticleModel{
		CreatedAt:    now,
		UpdatedAt:    now,
		Title:        cr.Title,
		Abstract:     cr.Abstract,
		Content:      cr.Content,
		UserID:       userID,
		UserNickName: userNickName,
		UserAvatar:   avatar,
		Category:     cr.Category,
		Source:       cr.Source,
		Link:         cr.Link,
		BannerID:     cr.BannerID,
		BannerUrl:    coverUrl,
		Tags:         cr.Tags,
	}

	err = article.Create()
	if err != nil {
		log.Errorw("创建文章失败", "err", err)
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OKWithMsg("文章发布成功", c)
}
