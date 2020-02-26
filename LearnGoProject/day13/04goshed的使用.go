package main

import (
	"fmt"
//	"time"
	"runtime"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("subtask go")
		}
	}()

	for i := 0; i < 2; i++ {
		fmt.Println("main")
		runtime.Gosched() //让出时间片，先让别的协程执行，它执行完，最后执行此协程
		//time.Sleep(time.Second)
	}
}
