package routers

import (
	v2 "gbv2/api/v2"
	"gbv2/middleware"
)

func (router RouterGroup) MessageRouter() {
	messageApi := v2.ApiGroupApp.MessageApi
	router.POST("message", messageApi.MessageCreateView)
	router.GET("message_all", middleware.JwtAdmin(), messageApi.MessageAllListView)
	router.GET("message", middleware.JwtAuth(), messageApi.MessageListView)
	router.POST("message_record", middleware.JwtAuth(), messageApi.MessageRecordView)
}
