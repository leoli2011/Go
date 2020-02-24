package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}
type Student struct {
	Person
	id int
	addr string
}

func main() {
	//结构体的顺序初始化，把所有字段都要初始化
	var s1 Student = Student{Person{"mike", 'M', 20}, 1, "bj"}
	fmt.Println("s1=", s1)
	//自动类型推导
	s2 := Student{Person{"mike", 'M', 20}, 2, "bj"}
	//%+v显示更详细信息
	fmt.Printf("s2=%+v\n", s2)

	//指定类型初始化，没有初始化的字段按照默认类型赋值
	s3 := Student{id:3}
	fmt.Printf("s3=%+v\n", s3)

	s4 := Student{Person:Person{name:"zhangshan"}, id:1}
	fmt.Printf("s4=%+v\n", s4)

	//s5 := Student{"mike", 'M', 18, 5, "bj"} //error
}
