package images_api

import (
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/res"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (ImagesApi) ImageRemoveView(c *gin.Context) {
	var re models.RemoveRequest
	err := c.ShouldBindJSON(&re)
	if err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}

	var imageList []models.ImageModel
	count := mysql.DB.Find(&imageList, re.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("文件不存在", c)
		return
	}
	mysql.DB.Delete(&imageList)
	total := strconv.Itoa(int(count))
	res.OKWithMsg("共删除"+total+"图片", c)
}
