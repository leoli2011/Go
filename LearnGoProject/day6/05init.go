package main

import "fmt"

//不需要用户来调用 系统自己调用 在main之前
func init()  {
	fmt.Println("init 做初始化操作")
}

func main() {
	fmt.Println("main")
}