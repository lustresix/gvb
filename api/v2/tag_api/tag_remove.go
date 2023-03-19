package tag_api

import (
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/res"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (TagApi) TagRemoveView(c *gin.Context) {
	var re models.RemoveRequest
	err := c.ShouldBindJSON(&re)
	if err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}

	var tagList []models.TagModel
	count := mysql.DB.Find(&tagList, re.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("标签不存在", c)
		return
	}
	mysql.DB.Delete(&tagList)
	total := strconv.Itoa(int(count))
	res.OKWithMsg("共删除"+total+"标签", c)
}
