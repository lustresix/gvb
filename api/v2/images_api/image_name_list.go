package images_api

import (
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/res"
	"github.com/gin-gonic/gin"
)

type ImageRep struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

func (ImagesApi) ImageNameListView(c *gin.Context) {
	var imageList []ImageRep

	mysql.DB.Model(models.ImageModel{}).Select("id", "path", "name").Scan(&imageList)

	res.OKWithData(imageList, c)

}
