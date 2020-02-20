package main

import "fmt"

func main() {

	//defer作为延迟调用，main函数结束之前调用
	//多个defer使用栈的方式排列
	defer fmt.Println("aaaaaaaaaa")
	defer fmt.Println("1111111111")
	defer fmt.Println("2222222222")
	fmt.Println("bbbbbbbbbb")
	fmt.Println("cccccccccc")
}
