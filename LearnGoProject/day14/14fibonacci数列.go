package main

import "fmt"

func fibonacci(ch chan<- int, flag <-chan bool)  {
	x, y := 1,1
	for {
		select {
		case ch<-x:
			x, y = y, x+y
		case f := <-flag:
			if f == true {
				fmt.Println("got the return flag")
				return
			}
		}
	}
}

func main() {
	ch := make(chan int)  //数字通信
	flag := make(chan bool) //程序是否退出
	go func() {
		for i:=0; i<8; i++ {
			num := <-ch
			fmt.Println("num", num)
		}
		flag<-true
	}()

	fibonacci(ch, flag)
}
