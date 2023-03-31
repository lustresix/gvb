package routers

import (
	v2 "gbv2/api/v2"
	"gbv2/middleware"
)

func (router RouterGroup) ArticleRouter() {
	ArticleApi := v2.ApiGroupApp.ArticleApi
	// 基本增删改查
	router.POST("article", middleware.JwtAuth(), ArticleApi.ArticleCreateView)
	router.GET("article", ArticleApi.ArticleListView)
	router.DELETE("article", middleware.JwtAuth(), ArticleApi.ArticleRemoveView)
	router.GET("article/detail", ArticleApi.ArticleDetailByTitleView)
	// 按日历获取
	router.GET("article/calendar", ArticleApi.ArticleCalendarView)
	// 标签
	router.GET("article/tag", ArticleApi.ArticleTagListView)
	// 文章详情
	router.GET("article/:id", ArticleApi.ArticleDetailView)
	// 收藏
	router.POST("article_collect", middleware.JwtAuth(), ArticleApi.ArticleCollCreateView)
	router.GET("article_collect", middleware.JwtAuth(), ArticleApi.ArticleListView)
	router.DELETE("article_collect", middleware.JwtAuth(), ArticleApi.ArticleCollCreateView)
}
