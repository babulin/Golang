package routers

import (
	"Golang/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	//基础路由
	BaseRouter()

	//User
	//beego.Router("/user", &controllers.User{})
	//beego.Router("/user/get-list", &controllers.User{}, "Get:GetList;Post:PostList")
	beego.AutoRouter(&controllers.User{})
}

func BaseRouter() {
	beego.Get("/", func(ctx *context.Context) {
		ctx.WriteString("GET:")
		ctx.Output.Body([]byte("嗨！请访问正常的路由"))
	})

	beego.Post("/", func(ctx *context.Context) {
		ctx.WriteString("POST:")
		ctx.Output.Body([]byte("嗨！请访问正常的路由"))
	})

	beego.Any("/", func(ctx *context.Context) {
		ctx.WriteString("Any:")
		ctx.Output.Body([]byte("嗨！请访问正常的路由"))
	})
}
