package main

import "fmt"

type Student2 struct {
	id int
	name string
	sex byte
	age int
	addr string
}

func main() {
	var s Student2
	var p1 *Student2
	p1 = &s

	p1.id = 1
	(*p1).name = "jack"
	p1.sex = 'M'
	p1.age = 18
	p1.addr = "bj"
	fmt.Println("s=",s)
	fmt.Println("p1=",p1)

	p2 := new(Student2)
	p2.id = 2
	p2.name = "mike"
	p2.sex = 'F'
	p2.age = 18
	p2.addr = "bj"
	fmt.Println("p2=", p2)

}
