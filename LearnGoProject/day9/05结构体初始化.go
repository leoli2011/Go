package main

import "fmt"

type student struct {
	id int
	name string
	sex byte
	age int
	addr string
}

func main()  {
	//顺序初始化，每个成员必须初始化
	var s1 student = student{1, "mike", 'M', 18, "bj"}
	fmt.Println("s1=", s1)

	//指定成员初始化，没有初始化的成员自动赋值0
	var s2 student = student {name:"mike", addr:"bj"}
	fmt.Println("s2=", s2)

	fmt.Println("---------------------")
	//定义一个结构体普通变量
	var s student
	s.id = 1
	fmt.Println("s=", s)
	//fmt.Printf("id=%d", s.id)
}
