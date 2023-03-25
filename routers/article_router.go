package routers

import (
	v2 "gbv2/api/v2"
	"gbv2/middleware"
)

func (router RouterGroup) ArticleRouter() {
	ArticleApi := v2.ApiGroupApp.ArticleModel
	router.POST("article", middleware.JwtAuth(), ArticleApi.ArticleCreateView)
	router.GET("article", ArticleApi.ArticleListView)
	router.GET("article/:id", ArticleApi.ArticleDetailView)
	router.GET("article/detail", ArticleApi.ArticleDetailByTitleView)
}
