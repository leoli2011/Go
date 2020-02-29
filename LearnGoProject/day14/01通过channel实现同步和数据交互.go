package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("主协程结束")
	ch := make(chan string)
	go func() {
		defer fmt.Println("子协程结束")
		for i := 0; i < 2; i++ {
			fmt.Println("子协程=", i)
			time.Sleep(time.Second)
		}
		ch <- "我是子协程， 子协程工作完毕"
	}()

	str := <-ch
	fmt.Println("str=", str)
}
