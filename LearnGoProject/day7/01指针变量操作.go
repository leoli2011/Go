package main

import "fmt"

func main() {
	var a int = 10

	fmt.Printf("a=%d\n", a)  //变量的内存
	fmt.Printf("&a=%v\n", &a) //变量的地址

	//保存地址必须要用指针(
	var p *int = &a
	fmt.Printf("p=%v, *p=%d\n", p, *p)
	*p = 100
	fmt.Printf("p=%v, *p=%d, a=%d\n", p, *p, a)

}
