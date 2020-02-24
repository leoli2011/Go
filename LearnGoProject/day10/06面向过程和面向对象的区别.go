package main

import "fmt"

func Add(a, b int) int  {
	return a+b
}

// 面向对象　　方法：给一个类型绑定一个函数
type long int
//给普通类型添加方法
func (a long)add01(b long) long {
	return a+b;
}

func main() {
	//考驾照---->买车----->自己开车---->亲自操作
	//　出租车　　　出租车.开车方法(目的地)

	var result int
	result = Add(1, 2)
	fmt.Printf("result=%d\n", result)

	var a long = 2
	//调用方式　变量.函数(所需要参数)
	result1 := a.add01(3)
	fmt.Println("result1=", result1)

}
