package main

import "fmt"

type Student3 struct {
	id int
	name string
	sex byte
	age int
	addr string
}

func main() {
	s1 := Student3{1, "mike", 'M', 18, "bj"}
	s2 := Student3{1, "mike", 'M', 18, "bj"}
	s3 := Student3{2, "mike", 'M', 18, "bj"}
	fmt.Println("s1==s2", s1==s2)
	fmt.Println("s1==s3", s1==s3)

	//同类型的两个结构体变量是可以直接赋值的
	var tmp  Student3
	tmp  = s3
	fmt.Println("tmp=", tmp)


	}