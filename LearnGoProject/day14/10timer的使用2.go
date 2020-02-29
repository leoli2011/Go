package main

import (
	"time"
	"fmt"
)

func main() {
	//验证timer时间到了，只会响应一次
	timer := time.NewTimer(3 * time.Second)
	for {
		<-timer.C
		fmt.Println("时间到")
	}
}
