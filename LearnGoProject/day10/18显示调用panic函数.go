package main

import (
	"fmt"
)

func testa()  {
	fmt.Println("aaaaaaaaaa")
}
func testb() {
	fmt.Println("bbbbbbbbbbbb")
	//显示调用panic的时候会导致程序中断
	panic("this is a panic test")
/*	a := 10
	b := 0
	c := a/b //当分母为0的时候，产生panic导致程序崩溃
	fmt.Println("c=",c)*/
}
func testc() {
	fmt.Println("ccccccccccc")
}
func main() {
	testa()
	testb()
	testc()
}
