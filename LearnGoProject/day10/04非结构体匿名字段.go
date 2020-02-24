package main

import "fmt"
type mystr string //自定义类型mystr
type Person3 struct {
	name string
	sex byte
	age int
}
type Student3 struct {
	Person3
	int
	mystr
}

func main() {
	s := Student3{Person3{"tom", 'M', 18}, 666, "zhangsan"}
	fmt.Printf("s=%+v\n", s)
	fmt.Println(s.name, s.sex, s.age, s.int, s.mystr)
	fmt.Println(s.Person3, s.int, s.mystr)
}
