package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个管道
	ch := make(chan string)
	go func() {
		defer fmt.Println("子协程调用完毕")
		for i:=0;i<2; i++  {
			fmt.Println("i=",i)
			time.Sleep(time.Second)
		}
		ch <-"子协程工作完毕"
	}()

	str := <-ch
	fmt.Println("str=", str)
}
