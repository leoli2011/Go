package main

import (
	"fmt"
)

func main() {
//	a := [5]int {24, 69, 80, 57, 13}
	a := [8]int {24, 69, 80, 57, 13, 100, 1 , 20}
	n := len(a)

	//升序
/*	for i := 0; i < n-1; i++ {
		for j := 0; j< n-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}*/
	//降序
	for i := 0; i < n-1; i++ {
		for j := 0; j< n-1-i; j++ {
			if a[j] < a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println("a=", a)
	fmt.Printf("排序后：\n")

	for i := 0; i < n; i++ {
		fmt.Printf("%d\n", a[i])
	}
	fmt.Println()
}
