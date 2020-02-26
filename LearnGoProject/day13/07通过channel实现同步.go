package main

import (
"fmt"
"time"
)

var ch = make(chan int)

//定义一个打印机，参数字符串，按照每个字符打印
func Printer1(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Println()
}

func Person11() {
	Printer1("Hello")
	ch<-666 //给管道写数据，发送
}

func Person22() {
	<- ch //从管道里边取数据， 接收，如果通道里面没有数据就会阻塞
	Printer1("World")

}

func main() {
	//两个协程共有一个资源
	go Person11()
	go Person22()
	for{
		time.Sleep(time.Second)
	}
}

