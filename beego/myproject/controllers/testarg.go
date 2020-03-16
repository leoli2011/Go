package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type TestargController struct {
	beego.Controller
}

type User struct {
	Id  int `form:"-"`
	Username string `form:"username"`
	Age  string `form:"age"`
	Email string `form:"email"`
}
/*

func (c *TestargController) Get() {
	id := c.GetString("id")
	c.Ctx.WriteString(id)

	name := c.Input().Get("name")
	c.Ctx.WriteString(name)
}
*/
func (c *TestargController) Get() {
	c.TplName = "index.tpl"
}

func (c *TestargController) Post() {
	u := User{}
	if err:= c.ParseForm(&u); err != nil {
		fmt.Printf("err=%s\n", err)

	}
	c.Ctx.WriteString("Username:="+u.Username+"Age:"+u.Age+"Email:"+u.Email)
}

