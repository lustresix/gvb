package comment_api

import (
	"fmt"
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/res"
	"gbv2/service/redis_ser"
	"github.com/gin-gonic/gin"
)

func (CommentApi) CommentDigg(c *gin.Context) {
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

	err = redis_ser.NewCommentDigg().Set(fmt.Sprintf("%d", cr.ID))
	if err != nil {
		res.FailWithMsg("评论失败", c)
		return
	}

	res.OKWithMsg("评论点赞成功", c)
	return

}
