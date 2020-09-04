package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id          int    `json:"id"`
	Nickname    string `json:"nickname"`
	Openid      string `json:"openid"`
	Mobile      string `json:"mobile"`
	Create_time string `json:"create_time"`
}

func init() {
	orm.RegisterModel(new(User))
}
