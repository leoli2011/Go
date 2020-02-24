package main

import (
	"fmt"
)
func xxx(arg ...interface{}) { //interface{} 是空接口 可以接收任意类型的值
	for i, j := range arg {
		fmt.Println(i, j)
	}

}

func main() {
	//空接口是万能类型，　保存任意类型的值
	var i interface{} = 1
	fmt.Println("i=", i)

	i = "abc"
	fmt.Println("i=", i)

	xxx(i)

}
