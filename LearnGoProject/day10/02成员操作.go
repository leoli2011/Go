package main

import "fmt"

type Person1 struct {
	name string
	sex byte
	age int
}
type Student1 struct {
	Person1
	id int
	addr string
}

func main() {
	s1 := Student1{Person1{"mike", 'M', 29}, 2, "bj"}
	//对象成员的操作方式
	s1.name = "yoyo"
	s1.sex = 'F'
	s1.age = 22
	s1.id = 7777
	s1.addr = "sh"
	//对象匿名成员操作方式２
	s1.Person1 = Person1{"Tom", 'm', 19}

	fmt.Println(s1.Person1.name)
	fmt.Println(s1.name) //匿名字段　s1.Person1.name == s1.name
	fmt.Println("s1=", s1)
}
