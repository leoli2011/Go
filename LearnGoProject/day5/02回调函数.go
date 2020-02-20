package main

import "fmt"

type FuncType01 func(int, int) int
//回调函数 就是有一个参数是函数类型， 这个函数就是回调函数
//ftest FuncType01 就相当于一个多态接口 调用一个接口可以有不同的实现方式

func calc(a, b int, ftest FuncType01) (result int) { //ftest = Add01
	result = ftest(a, b)
	// 相当于 result = Add01(a, b)
	return
}

func Add01(a, b int) int {
	return a+b
}

func Minus01(a, b int) int {
	return a-b
}

func main() {
	result :=calc(3, 15, Add01)
	fmt.Println("result=", result)

	result = calc(10, 5, Minus01)
	fmt.Println("result=", result)
}
