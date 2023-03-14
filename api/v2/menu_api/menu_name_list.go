package menu_api

import (
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/res"
	"github.com/gin-gonic/gin"
)

type MenuNameResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Path  string `json:"path"`
}

func (MenuApi) MenuNameList(c *gin.Context) {
	var menuNameList []MenuNameResponse
	mysql.DB.Model(models.MenuModel{}).Select("id", "title", "path").Scan(&menuNameList)
	res.OKWithData(menuNameList, c)
}
