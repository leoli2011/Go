package main

import "fmt"

//通过type给一个函数类型命名   函数也是一种数据类型(类似于 int byte)
type FuncType func(a, b int) int //没有名字的函数  没有{}

func Add(a, b int) int {
	return a+b
}

func xxxx(c, d int) int {
	return c-d
}

func main() {
	var a int
	a = 10
	fmt.Println("a=",a)

	var ftest FuncType;
	ftest = Add
	result := ftest(3,5) //等价于 Add(3,5)
	fmt.Println("result=", result)

	ftest = xxxx
	result = ftest(3, 2) //等价于XXXX(3,2)
	fmt.Println("result=", result)
}
