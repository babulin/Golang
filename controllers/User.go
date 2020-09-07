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
func (c *User) Insert() {
	//数据库..
	db := orm.NewOrm()
	user := models.User{Nickname: "李一鸣", Openid: "zxgsdweeffsd", Mobile: "15736107111", Create_time: "2020-09-07 12:00:00"}

	id, err := db.Insert(&user)
	fmt.Println(id, err)
	arr := []int{1, 2, 3, 4, 5}

	msg := fmt.Sprintf("保存成功%d", id)
	c.Json(&Json{0, msg, arr})
}

//指定Post访问方法
func (c *User) PostList() {
	c.Ctx.WriteString("PostList")
}
