package controllers

import (
	"Golang/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Controller
}

func (c *User) Prepare() {
	//c.Msg(1,"公共验证")
}

//访问Get路由
func (c *User) Get() {

	//数据库..
	db := orm.NewOrm()
	var sUser = models.User{Id: 1}
	db.Read(&sUser)
	c.Json(&Json{1, "消息", sUser})
}

//访问Post路由
func (c *User) Post() {
}

//指定Get访问方法
func (c *User) GetList() {
	var str string
	str = fmt.Sprintf("%d", PAGE_NUM)
	c.Ctx.WriteString("GetList" + str)
}

//指定Post访问方法
func (c *User) PostList() {
	c.Ctx.WriteString("PostList")
}
