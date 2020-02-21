package main

import "fmt"

func main() {
	a := [5]int {1, 2,3, 4,5}
	b := [5]int {1,2,3,4,5}
	c := [5]int {1,2,3}
	fmt.Println("a==b", a==b)
	fmt.Println("a==c", a==c)

	//同类型的数组可以直接赋值
	var d [5]int
	d = a
	fmt.Println("d=", d)
	d = c;
	fmt.Println("d=", d)

}
