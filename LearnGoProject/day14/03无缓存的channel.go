package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个无缓存区的channel
	ch := make(chan int, 0)
	//len(ch)缓冲区使用数据个数， cap(ch)缓冲区大小
	fmt.Printf("len(%d), cap(%d)\n", len(ch), cap(ch))

	go func() {
		for i:=0; i<3; i++ {
			fmt.Printf("子协程 i=%d\n", i)
			ch <-i
		}
	}()

	time.Sleep(2*time.Second)

	for i := 0; i < 3; i++ {
		num := <-ch
		fmt.Printf("主协程i=%d\n", num)
	}
}
