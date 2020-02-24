package main

import (
	"fmt"
)

func testa1()  {
	fmt.Println("aaaaaaaaaa")
}
func testb1(x int) {
	defer func() {
		//recover() //通过recover, 捕获错误性信息，继续往后执行
		//fmt.Println(recover())
		if err:=recover(); err != nil { //打印出产生的异常
			fmt.Println(err)
		}
	}() //别忘了调用次函数
	a := [10]int{0}
	//var a [10]int
	a[x] = 1111 //当x=20时，数组越界，产生一个panic, 导致程序崩溃
}
func testc1() {
	fmt.Println("ccccccccccc")
}
func main() {
	testa1()
	testb1(20)
	testc1()
}
