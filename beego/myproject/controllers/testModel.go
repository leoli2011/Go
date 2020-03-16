package controllers

import (
	"github.com/astaxie/beego" //beego的web端
	"myproject/models"
)

type TestModelController struct {
	beego.Controller
}


func (c *TestModelController) Get()  {

	user := models.UserInfo{Id:80, Username:"wangwu", Password:"20203399"}
	models.AddUser(&user)
	c.Ctx.WriteString("call model success")
}
