package main

import "fmt"

func main() {
	//切片与数组的区别
	a := [5]int {}
	fmt.Printf("len=%d, cap=%d \n", len(a), cap(a))

	s := []int {}
	fmt.Printf("len=%d, cap=%d \n", len(s), cap(s))
	fmt.Printf("s type=%T \n", s);

	s = append(s, 1)
	fmt.Printf("len=%d, cap=%d \n", len(s), cap(s))

	s = append(s, 1)
	fmt.Printf("len=%d, cap=%d \n", len(s), cap(s))

	s = append(s, 1)
	fmt.Printf("len=%d, cap=%d \n", len(s), cap(s))

	s = append(s, 1)
	fmt.Printf("len=%d, cap=%d \n", len(s), cap(s))

	s = append(s, 1)
	fmt.Printf("len=%d, cap=%d \n", len(s), cap(s))

}
