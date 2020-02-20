package main

import "fmt"

func test() func() int {
	var x int //x没有初始化， 值为0
	return func() int {
		x++
		return x*x
	}
}
func main() {
	//返回值是一个匿名函数，返回一个函数类型，通过F来调用返回的匿名函数，F来调用闭包函数
	//F不关心这些捕获的变量和常量是否超出作用域，只要你的F存在，还在使用X，这个变量X就会一直存在
	//只要闭包在使用X，这些变量就一直存在
	f :=test()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	f = test()
	fmt.Println(f())
}
/*
func test() int {
	var x int //没有初始化值为0
	x++
	return x
}

func main() {
	fmt.Println(test())
	fmt.Println(test())
	fmt.Println(test())
	fmt.Println(test())
}
*/