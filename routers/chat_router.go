package routers

import (
	v2 "gbv2/api/v2"
)

func (router RouterGroup) ChatRouter() {
	chatApi := v2.ApiGroupApp.ChatApi
	router.GET("chat", chatApi.CommentRemoveView)
	router.GET("chat_group", chatApi.ChatGroupView)
}
