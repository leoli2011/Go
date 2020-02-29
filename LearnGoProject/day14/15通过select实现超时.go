package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int) //数据
	quit := make(chan bool)

	go func() {
		for {
			select {
			case data := <-ch:
				fmt.Println("data=", data)
			case <-time.After(time.Second*10):
				quit<-true
				fmt.Println("子协程： 结束")
			}
		}
	}()

	for i:=0; i<5; i++ {
		ch<-i
		time.Sleep(time.Second)
	}
	flag := <-quit
	fmt.Printf("程序结束, %v", flag)
}
