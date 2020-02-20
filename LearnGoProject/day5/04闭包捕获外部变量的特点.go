package main

import "fmt"

func main() {
	a := 10
	b := "aaa"

	func() {
		//闭包以引用的方式捕获外部变量
		a = 40
		b = "666"
		fmt.Printf("a=%d, b =%s\n", a, b)
	}()
	fmt.Printf("a=%d, b =%s\n", a, b)
}
