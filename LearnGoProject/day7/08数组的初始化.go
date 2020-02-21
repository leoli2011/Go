package main

import "fmt"

func main() {
	//声明定义同时初始化

	var a [5]int = [5]int{1, 2,3 ,4 ,5}
	fmt.Println("a=", a)

	//自动推倒
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b=", b)

	//部分初始化
	c := [5]int{1,2,3}
	fmt.Println("c=", c)

	//某个元素的初始化
	d := [5]int{2:10, 4:20}
	fmt.Println("d=", d)

}
