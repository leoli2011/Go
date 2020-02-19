package main

import "fmt"

func Myfunc111(temp ...int) {
	for _,data := range temp {
		fmt.Println("data=", data)
	}
}
func test(args ...int) {
//	Myfunc111(args...) //实参一定是args...
//	Myfunc111(args[0:2]...) //从0开始到2不包含2传递过去
//	Myfunc111(args[:3]...) //从0开始到2不包含2传递过去
	Myfunc111(args[2:]...)//从args[2]开始包含本身2到最后
}
func main() {
	test(1,2,3,4,5,6)
	
}
