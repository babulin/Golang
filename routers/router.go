package routers

import (
	"Golang/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/list", &controllers.UserController{}, "get:GetList;post:PostList")
}
