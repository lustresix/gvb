package routers

import v2 "gbv2/api/v2"

func (router RouterGroup) MenuRouter() {
	app := v2.ApiGroupApp.MenuApi
	router.POST("menu", app.MenuCreatView)
	router.GET("menu", app.MenuListView)
	router.GET("menu_name", app.MenuNameList)
	router.GET("menus/:id", app.MenuDetailView)
	router.PUT("menu/:id", app.MenuUpdateView)
	router.DELETE("menu", app.MenuRemoveView)

}
