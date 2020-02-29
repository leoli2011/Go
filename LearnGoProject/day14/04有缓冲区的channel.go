package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)

	fmt.Printf("len(ch)=%d, cap(ch)=%d\n", len(ch), cap(ch))

	go func() {
		for i:=0; i<10; i++ {
			ch <-i  //往chan写内容
			fmt.Printf("子协程[%d], len=%d, cap=%d\n", i, len(ch), cap(ch))
		}
	}()

	time.Sleep(3* time.Second)
	for i:=0; i<10; i++ {
		num := <-ch
		fmt.Printf("num=%d\n", num)
	}
}
