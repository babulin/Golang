package controllers

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Data["name"] = "张三"
	c.Data["msg"] = "这是UserController的msg"
	c.TplName = "user/index.tpl"
}

func (c *UserController) Post() {
	c.Ctx.WriteString("aaaaaaaaa")
}
func (c *UserController) GetList() {
	c.Ctx.WriteString("GetList")
}
func (c *UserController) PostList() {
	c.Ctx.WriteString("PostList" + beego.AppConfig.String("httpport"))
}
