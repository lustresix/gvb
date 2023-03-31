package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(viper.GetString("system.env"))
	router := gin.Default()
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	apiRouterGroup := router.Group("api")

	routerGroupApp := RouterGroup{apiRouterGroup}
	// 系统配置api
	routerGroupApp.SettingRouter()
	// 图片上传
	routerGroupApp.ImagesRouter()
	// 广告
	routerGroupApp.AdvertRouter()
	// 菜单
	routerGroupApp.MenuRouter()
	// 用户
	routerGroupApp.UserRouter()
	// 广告
	routerGroupApp.TagRouter()
	// 消息
	routerGroupApp.MessageRouter()
	// 文章
	routerGroupApp.ArticleRouter()
	// GPT
	routerGroupApp.GptRouter()
	// 点赞
	routerGroupApp.DiggRouter()

	return router
}
