package main

import "fmt"

func test(m1 map[int]string) { //map和切片一样是引用类型
	delete(m1, 1)
}
func main() {
	m1 := map[int]string{1:"jack", 2:"tony", 3:"marry"}
	//map输出不保证顺序，是无序的　
	fmt.Println("m1=", m1)

	test(m1)
	fmt.Println("m1=",m1)

}
