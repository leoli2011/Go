package main

import "fmt"

func test1(a int) {
	if a==1 { //终止条件
		fmt.Println("a=", a)
		return
	}//函数调用本身
	test1(a-1)
	fmt.Println("a=", a)
}

func main() {
	test1(3)
	fmt.Println("main")
}
