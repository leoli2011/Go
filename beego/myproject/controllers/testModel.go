package controllers

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"  //操作mysql的
	"github.com/astaxie/beego" //beego的web端
	"fmt"
) //驱动

//UserInfo === 数据库里的user_info
type UserInfo struct {
	Id int64
	Username string
	Password string
}

type TestModelController struct {
	beego.Controller
}

func init()  {
	err := orm.RegisterDataBase("default", "mysql", "root:root@/test?charset=utf8")
	//err := orm.RegisterDataBase("default", "mysql", "root:root@tcp(123.63.17.199:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	//register model
	orm.RegisterModel(new(UserInfo)) //创建一个user_info表
}

func (c *TestModelController) Get()  {

	//插入操作
	o := orm.NewOrm()
	user := UserInfo{Id:10, Username:"zhangsan", Password:"123456"}
	id, err := o.Insert(&user)

	c.Ctx.WriteString(fmt.Sprintf("操作数据库ID=%d, err=%v", id, err))

	user = UserInfo{Id:20, Username:"lisi", Password:"654321"}
	id, err = o.Insert(&user)
	c.Ctx.WriteString(fmt.Sprintf("操作数据库ID=%d, err=%v", id, err))
	//更新操作
	user = UserInfo{10, "hahaha", "8888"}
	o = orm.NewOrm()
	o.Update(&user)
	c.Ctx.WriteString(fmt.Sprintf("操作数据库ID=%d, err: %v", id, err))

	//读取操作
	user = UserInfo{Id:20}
	o.Read(&user)
	c.Ctx.WriteString(fmt.Sprintf("读取数据库user=%v", user))

}
