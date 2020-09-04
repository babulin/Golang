//+----------------------------------
//| 公共配置区域
//+----------------------------------

package controllers

//常量
const PAGE_NUM = 12

//json结构图
type Json struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//公共接口
type IController interface {
	Json(json *Json)
	Msg(code int, msg string, data string)
}
