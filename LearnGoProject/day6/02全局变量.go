package main

import "fmt"

// 定义在花括号之外的叫全局变量,
//程序消失才消失
func test() {
	fmt.Println("test a=",a)
}

var a int

func main() {
	a = 10
	fmt.Println("a=",a)
	test();
}
