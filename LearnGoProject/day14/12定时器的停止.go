package main

import (
	"time"
	"fmt"
)

func main() {
	timer := time.NewTimer(3*time.Second)
	ok := timer.Reset(5*time.Second) //把以前3秒重置为1秒
	fmt.Println("ok=", ok)
	<-timer.C
	fmt.Println("时间到")
}
/*
func main() {
	timer := time.NewTimer(3 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("子协程 时间到了")
	}()

	timer.Stop()
	for {

	}
}
*/
