package main

import (
	"fmt"
	"runtime"
	"time"
)

func test() {
	defer 	fmt.Println("bbbbb")
	//return //终止函数
	runtime.Goexit() //终止所有子协程
	fmt.Println("ddddddd")
}
func main() {
	go func() {
		fmt.Println("aaaaaaa")
		test()
		fmt.Println("cccccccc")
	}()
	for {
		fmt.Println("this is main")
		time.Sleep(time.Second)
	}
}
