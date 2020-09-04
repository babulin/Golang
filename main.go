package main

import (
	_ "Golang/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	//连接数据库
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("database"))

	//自动渲染模板
	beego.BConfig.WebConfig.AutoRender = false
	beego.Run()
}
