package routers

import (
	v2 "gbv2/api/v2"
)

func (router RouterGroup) TagRouter() {
	tagApi := v2.ApiGroupApp.TagApi
	router.POST("tag", tagApi.TagCreateView)
	router.GET("tag", tagApi.TagListView)
	router.PUT("tag/:id", tagApi.TagUpdateView)
	router.DELETE("tag", tagApi.TagRemoveView)
}
