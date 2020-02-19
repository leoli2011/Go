package main

import "fmt"

//无参数只有一个返回值，通过return返回
func Myfunc01() int {
	return 5555;
}

//这种返回值方式是官方推荐
func Myfunc02() (a int) {
	return 6666;
}

//常用写法
func Myfunc03() (a int) {
	a = 888
	return
}
func main() {
	var a int
	a = Myfunc01()
	fmt.Println(a)
	b := Myfunc01();
	fmt.Printf("b=%d\n", b)
	c := Myfunc03();
	fmt.Printf("c=%d", c)
}
