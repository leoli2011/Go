package main

import "fmt"

type Person2 struct {
	name string
	sex byte
	age int
}
type Student2 struct {
	Person2
	id int
	addr string
	name string
}

func main() {
	var s Student2
	//就近原则　如果能在本作用域找到此成员，就操作此成员，如果找不到，就找到继承的字段
	s.name = "zhangsan"//赋值的是Student2的name，就近原则
	s.sex = 'M'
	s.age = 18
	s.addr = "bj"
	fmt.Printf("s=%+v", s)
}
