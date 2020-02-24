package main

import "fmt"

type Person5 struct {
	name string
	sex byte
	age int
}
type Student5 struct {
	*Person5
	id int
	addr string
}

func main() {
	s1 := Student5{&Person5{"mike", 'M', 18}, 777, "bj"}
	fmt.Printf("s1=%+v\n", s1)
	fmt.Println(s1.Person5.name)
	fmt.Println(s1.name)

	var s2 Student5
	s2.Person5 = new(Person5)
	s2.name = "joke"
	s2.sex = 'M'
	s2.age = 18
	s2.id = 2222
	s2.addr = "bj"
	//fmt.Printf("%+v", s2)
	fmt.Println(s2.name,s2.sex,s2.age, s2.id, s2.addr)

}
