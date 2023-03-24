package routers

import (
	v2 "gbv2/api/v2"
	"gbv2/middleware"
)

func (router RouterGroup) ArticleRouter() {
	ArticleApi := v2.ApiGroupApp.ArticleModel
	router.POST("article", middleware.JwtAuth(), ArticleApi.ArticleCreateView)
}
