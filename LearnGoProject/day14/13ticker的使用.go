package main

import (
	"time"
	"fmt"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)
	i:=0
	for {
		i++
		<-ticker.C //两秒钟的间隔
		fmt.Println("i=",i)
		if i==5 {
			ticker.Stop()
			break
		}
	}
}
