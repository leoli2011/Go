package main

import "fmt"

func main() {
	a := [10]int {0,1,2,3,4,5,6,7,8,9}

	s1 := a[2:5]
	s1[1] = 666
	fmt.Println("s1=", s1)
	fmt.Println("a=", a)
	//容量没变都引用同一个数组
	s2 := s1[2:7]
	s2[1] = 777
	fmt.Println(s2)
	fmt.Println(a)

}
