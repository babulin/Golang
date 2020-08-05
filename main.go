package main

import (
	_ "Golang/routers"
	"github.com/astaxie/beego"
)

func main() {
	//自动渲染模板
	beego.BConfig.WebConfig.AutoRender = false
	beego.Run()
}
