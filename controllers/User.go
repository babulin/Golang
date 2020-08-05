package controllers

import "github.com/astaxie/beego"

type User struct {
	beego.Controller
}

func (c *User) Get() {
	c.Data["name"] = "张三"
	c.Data["msg"] = "这是UserController的msg"
	c.TplName = "user/index.tpl"
}

func (c *User) Post() {
	c.Ctx.WriteString("aaaaaaaaa")
}
func (c *User) GetList() {
	c.Ctx.WriteString("GetList")
}
func (c *User) PostList() {
	c.Ctx.WriteString("PostList" + beego.AppConfig.String("httpport"))
}
