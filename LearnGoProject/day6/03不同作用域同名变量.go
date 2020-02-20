package main

import "fmt"

var aa byte  //全局变量

func main() {
	var aa int //局部变量
	fmt.Printf("a=%T\n", aa)
	{
		var aaa float32
		fmt.Printf("444=%T\n", aaa)
	}
}
