package routers

import (
	v2 "gbv2/api/v2"
)

func (router RouterGroup) TagRouter() {
	tagApi := v2.ApiGroupApp.TagApi
	router.POST("Tags", tagApi.TagCreateView)
	router.GET("Tags", tagApi.TagListView)
	router.PUT("Tags/:id", tagApi.TagUpdateView)
	router.DELETE("Tags", tagApi.TagRemoveView)
}
