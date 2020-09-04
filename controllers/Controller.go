package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type Controller struct {
	beego.Controller
}

func (c *Controller) Json(json *Json) {
	c.Data["json"] = json
	c.ServeJSON()
}

func (c *Controller) Msg(code int, msg string) {
	var str string
	str = fmt.Sprintf("{\"code\":\"%d\",msg:\"%v\"}", code, msg)
	c.Ctx.WriteString(str)
}
