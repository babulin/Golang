package controllers

import "github.com/astaxie/beego"

type Api struct {
	beego.Controller
}

//访问Get
func (_this *Api) Get() {
	_this.Ctx.WriteString("访问了Get方法")
}

func (_this *Api) Post() {
	_this.Ctx.WriteString("访问了Post方法")
}

func (_this *Api) GetList() {
	_this.Ctx.WriteString("访问了GetList方法")
}
