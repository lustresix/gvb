package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(viper.GetString("system.env"))
	router := gin.Default()

	apiRouterGroup := router.Group("api")

	routerGroupApp := RouterGroup{apiRouterGroup}
	// 系统配置api
	routerGroupApp.SettingRouter()
	return router
}
