package main

import "fmt"

type Person6 struct {
	name string
	sex byte
	age int
}

//Person6类型，实现了一个类型
func (p Person6) PrintInfo() {
	fmt.Printf("name=%s, sex=%c, age=%d", p.name, p.sex, p.age)
}

//有个学生，　继承了Person字段，　成员和方法都继承了
type Student6 struct {
	Person6
	id int
	addr string
}

//方法的重写　Student6实现了一个方法，这个方法和Person方法同名，这种方法叫重写
func (tmp Student6)PrintInfo() {
	fmt.Println("Student6=", tmp)
}

func main() {
	s6 := Student6{Person6{"tom", 'M', 18}, 10, "bj"}
	//就近原则：先找本作用域的方法，找不到再去找继承的方法
	//如果找不到，调用报错
	s6.PrintInfo()
	//显示调用继承的方法
	s6.Person6.PrintInfo()
}
