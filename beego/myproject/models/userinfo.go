package models

import (
	_ "github.com/go-sql-driver/mysql" //驱动
	"github.com/astaxie/beego/orm"  //操作mysql的
)

//UserInfo === 数据库里的user_info
type UserInfo struct {
	Id int64
	Username string
	Password string
}

var db orm.Ormer

func init()  {
	orm.Debug = true
	err := orm.RegisterDataBase("default", "mysql", "root:root@/test?charset=utf8")
	//err := orm.RegisterDataBase("default", "mysql", "root:root@tcp(123.63.17.199:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	//register model
	orm.RegisterModel(new(UserInfo)) //创建一个user_info表
	db = orm.NewOrm()
}

func AddUser(user *UserInfo)(int64, error)  {
	id, err := db.Insert(user)
	return id, err
}

func ReadUser(users *[]UserInfo)  {
	qb,_ := orm.NewQueryBuilder("mysql")
	qb.Select("*").From("user_info")
	sql := qb.String()
	db.Raw(sql).QueryRow(users)
}