package main

import "fmt"

type Student4 struct {
	id int
	name string
	sex byte
	age int
	addr string
}

func modify01(s1 Student4)  {
	s1.id = 666
	fmt.Println("modify s1=", s1)
}

func modify02(s1 *Student4)  {
	s1.id = 888
	fmt.Println("modify s1=", s1)
}
func main() {
	s1 := Student4{1, "mike", 'M', 18, "bj"}
	modify01(s1) //值传递
	fmt.Println("main s1=", s1)
	modify02(&s1) //地址传递
	fmt.Println("main s1=", s1)
}
