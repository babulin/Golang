package routers

import (
	"Golang/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//home
	beego.AutoRouter(&controllers.Home{})

	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/list", &controllers.User{}, "get:GetList;post:PostList")
}
