package main

import (
	"time"
	"fmt"
)

func main() {
	<-time.After(3*time.Second) //定时3S，阻塞3S，3S后产生一个事件往channel里写内容，和第一种一样
	fmt.Println("时间到")

}

/*func main() {
	time.Sleep(3*time.Second)
	fmt.Println("时间到")
}*/

/*
func main() {
	timer := time.NewTimer(3*time.Second)
	<-timer.C
	fmt.Println("时间到")
}*/
