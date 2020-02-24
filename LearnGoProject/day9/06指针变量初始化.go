package main

import "fmt"

type Person struct {
	id int
	name string
	sex byte
	age int
	addr string
}

func main() {
	var p1 *Person = &Person{1, "mike", 'M', 18, "bj"}
	fmt.Printf("p1=%+v  \n", p1)

	p2 := &Person{name:"mike", addr:"bj"}
	fmt.Printf("p2 type=%T\n", p2 )
	fmt.Println("p2=", p2)
}
