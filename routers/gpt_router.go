package routers

import (
	v2 "gbv2/api/v2"
)

func (router RouterGroup) GptRouter() {
	app := v2.ApiGroupApp.GPTApi
	router.POST("gpt",app.Chat)
}

