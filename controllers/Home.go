package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type Home struct {
	beego.Controller
}

func (c *Home) Index() {

	//int [指定类型]
	var i int = 100

	//float [key := value] 根据值进行判断类型
	f := 123.05

	//定义字符串
	var str string = "World"

	fmt.Println("Console 输出语句")
	fmt.Printf("格式化:%d\n", i)
	fmt.Println("Hello" + " " + str)
	fmt.Printf("float:%.2f\n", f)
}

func (c *Home) Get() {

}
