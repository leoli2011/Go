package main

import "fmt"

func myfunc001() (int, int, int) {
	return 111,222,333
}

//go 官方推荐写法
func myfunc002() (a int, b int, c int) {
	a, b, c = 55,66,77
	return
}

func main() {
	a, b, c := myfunc001()
	fmt.Printf("a=%d,b=%d,c=%d\n", a, b, c)
	a, b, c = myfunc002()
	fmt.Printf("a=%d,b=%d,c=%d\n", a, b, c)
}
