package main

import "fmt"

func main() {
	var a int
	fmt.Printf("请输入变量a: \n")
	//fmt.Scan(&a)
	fmt.Scanf("%d", &a);
	fmt.Println("a=", a)
}
