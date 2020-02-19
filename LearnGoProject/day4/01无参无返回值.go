package main

import "fmt"

func myfunction() {
	b := "发射子弹"
	fmt.Println(b)
}

func main() {
	//无参数无返回值的参数调用：函数名()
	myfunction()
	myfunction()
}
