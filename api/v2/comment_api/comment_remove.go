package comment_api

import (
	"fmt"
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/res"
	"gbv2/service/redis_ser"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentIDRequest struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

func (CommentApi) CommentRemoveView(c *gin.Context) {
	var cr CommentIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}
	var commentModel models.CommentModel
	err = mysql.DB.Take(&commentModel, cr.ID).Error
	if err != nil {
		res.FailWithMsg("评论不存在", c)
		return
	}
	// 统计评论下的子评论数 再把自己算上去
	var subCommentList []models.CommentModel
	FindSubComment(commentModel, &subCommentList)
	count := len(subCommentList) + 1
	err = redis_ser.NewCommentCount().Set(commentModel.ArticleID)
	if err != nil {
		return
	}
	// 判断是否是子评论
	if commentModel.ParentCommentID != nil {
		mysql.DB.Model(&models.CommentModel{}).
			Where("id = ?", *commentModel.ParentCommentID).
			Update("comment_count", gorm.Expr("comment_count - ?", count))
	}
	// 删除子评论以及当前评论
	var deleteCommentIDList []uint
	for _, model := range subCommentList {
		deleteCommentIDList = append(deleteCommentIDList, model.ID)
	}
	for _, i := range deleteCommentIDList {
		mysql.DB.Model(models.CommentModel{}).Delete("id = ?", i)
	}

	res.OKWithMsg(fmt.Sprintf("删除 %d 个评论", len(deleteCommentIDList)), c)
}
