package routers

import v2 "gbv2/api/v2"

func (router RouterGroup) DiggRouter() {
	app := v2.ApiGroupApp.DiggApi
	router.POST("digg_article", app.DiggArticle)
}
