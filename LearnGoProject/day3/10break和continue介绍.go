package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for{
		i++
		time.Sleep(time.Second) //睡觉一秒钟
		if i==5 {
			//break//跳出本层循环
			continue   //跳出本次循环
		}
		fmt.Printf("i=%d \n", i)
	}

}
