package controllers

import (
	"github.com/astaxie/beego" //beego的web端
	"myproject/models"
)

type TestModelController struct {
	beego.Controller
}


func (c *TestModelController) Get()  {

	//user := models.UserInfo{Id:80, Username:"wangwu", Password:"20203399"}
	//models.AddUser(&user)
	//c.Ctx.WriteString("call model success")

	var users []models.UserInfo
	models.ReadUser(&users)
/*	for i := 0; i < len(users); i++ {
		c.Ctx.WriteString(fmt.Sprintf("tt users:%v", users[i]))
	}*/
	c.Data["Title"] = "leoli"
	c.Data["Users"] = users
	//fmt.Printf("users:%v", users)
	//c.Ctx.WriteString(fmt.Sprintf("users:%v", users))
	c.Data["len"] = len(users)
	c.TplName = "test.tpl"
}
