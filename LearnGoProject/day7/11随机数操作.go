package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main()  {
	//如果种子一样，那么每次运行程序产生的随机数都一样
	//每次运行让种子不一样才能程序每次运行产生不一样的随机数
	//rand.Seed(6666)

	rand.Seed(time.Now().UnixNano())//以当前系统时间作为种子
	for i := 0; i < 6; i++ {
		//fmt.Println("rand=", rand.Int())
		fmt.Println("rand=", rand.Intn(10)) //产生10之内的随机数
	}
}
