package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 0)
	fmt.Printf("len=%d, cap=%d\n",len(ch), cap(ch))

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("子协程: i=%d\n", i)
			ch <-i
		}
	}()

	time.Sleep(time.Second)
	for i := 0; i < 3; i++ {
		num:= <- ch
		fmt.Printf("num=%d\n", num)
	}
}
