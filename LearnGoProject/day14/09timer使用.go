package main

import (
	"time"
	"fmt"
)

func main() {
	//创建一个定时器，设置时间为3秒，3秒后，往time通道写内容
	timer := time.NewTimer(3*time.Second)
	fmt.Println("time is", time.Now())
	t := <-timer.C
	fmt.Println("t=", t)
}
