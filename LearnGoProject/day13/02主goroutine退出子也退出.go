package main

import (
	"fmt"
	"time"
)

func NewTask() {
	for {
		fmt.Println("this is  a newTask")
		time.Sleep(time.Second)
	}
}

func main() {
	go NewTask()     //go关键字就新建一个协程，相当于新建一个业务
/*	for {
		fmt.Println("this is main gorouting")
		time.Sleep(time.Second)
	}*/
	i :=0
	for {
		i++
		fmt.Println("this is main gorouting")
		time.Sleep(time.Second)
		if (i==2){
			break
		}
	}
}
