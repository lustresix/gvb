package comment_api

import (
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/res"
	"gbv2/service/es_ser"
	"gbv2/service/redis_ser"
	"gbv2/utils/jwt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentRequest struct {
	ArticleID       string `json:"article_id" binding:"required" msg:"请选择文章"`
	Content         string `json:"content" binding:"required" msg:"请输入评论内容"`
	ParentCommentID *uint  `json:"parent_comment_id"` // 父评论id
}

func (CommentApi) CommentCreateView(c *gin.Context) {
	var cr CommentRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	// 查找文章是否存在
	_, err = es_ser.CommDetail(cr.ArticleID)
	if err != nil {
		res.FailWithMsg("文章不存在", c)
	}
	// 判断是否是子评论
	if cr.ParentCommentID != nil {
		// 子评论
		var parentComment models.CommentModel
		err := mysql.DB.Take(&parentComment, cr.ParentCommentID).Error
		if err != nil {
			res.FailWithMsg("父评论不存在", c)
			return
		}
		// 父评论是否是当前文章的
		if parentComment.ArticleID != cr.ArticleID {
			res.FailWithMsg("文章错误", c)
			return
		}
		// 给父评论评论数+1
		mysql.DB.Model(&parentComment).Update("comment_count", gorm.Expr("comment_count + 1"))
	}
	// 添加评论
	mysql.DB.Create(&models.CommentModel{
		ParentCommentID: cr.ParentCommentID,
		Content:         cr.Content,
		ArticleID:       cr.ArticleID,
		UserID:          claims.UserID,
	})
	redis_ser.Comment(cr.ArticleID)
	res.OKWithMsg("文章评论成功", c)
	return
}
