package main

import "fmt"

func main() {
	//var a int
	//a = 10
	a := 5
	if (a==10) {
		fmt.Println("a==10")
	}

	//if语句初始化然后赋值
	if a:=10; a==10 {
		fmt.Println("a==10")
	} else {
		fmt.Println("a!=10")
	}

	fmt.Println("a=", a)

}
