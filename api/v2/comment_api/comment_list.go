package comment_api

import (
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/res"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

type CommentListRequest struct {
	ArticleID string `form:"article_id"`
}

func (CommentApi) CommentListView(c *gin.Context) {
	var cr CommentListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}
	list := FindArticleCommentList(cr.ArticleID)
	res.OKWithData(filter.Select("c", list), c)
	return
}

// FindArticleCommentList 获取评论列表
func FindArticleCommentList(articleID string) (RootCommentList []*models.CommentModel) {
	// 先找父评论 -- parent_comment_id is null
	mysql.DB.Preload("User").Find(&RootCommentList, "article_id = ? and parent_comment_id is null", articleID)
	for _, model := range RootCommentList {
		var subCommentList []models.CommentModel
		FindSubComment(*model, &subCommentList)
		model.SubComments = subCommentList
	}
	return
}

// FindSubComment 递归查评论下的子评论
func FindSubComment(model models.CommentModel, subCommentList *[]models.CommentModel) {
	mysql.DB.Preload("SubComments.User").Take(&model)
	for _, sub := range model.SubComments {
		*subCommentList = append(*subCommentList, sub)
		FindSubComment(sub, subCommentList)
	}
	return
}
