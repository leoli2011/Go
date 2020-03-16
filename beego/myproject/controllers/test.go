package controllers

import "github.com/astaxie/beego"

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	c.Ctx.WriteString("这是一个beego 控制器Get方法")
}

func (c *TestController) Post() {
	c.Ctx.WriteString("这是一个beego 控制器Post方法")
}
