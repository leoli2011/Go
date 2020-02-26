package main

import (
	"fmt"
	"time"
)

//主程序退出 子程序跟着退出 导致子协程没有被调用
func main() {
	go func() {
		i:=0
		for {
			i++
			fmt.Println("子协程 i=", i)
			time.Sleep(time.Second)
		}
	}()

/*	i := 0
	for{
		i++
		fmt.Println("main i=",i)
		time.Sleep(time.Second)
		if i == 2{
			break
		}
	}*/
}
