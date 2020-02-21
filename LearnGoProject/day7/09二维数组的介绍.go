package main

import (
	"fmt"
)

func main() {
	var a[3][4] int
	k := 0
	for i := 0; i < 3; i++ {
		for j:=0; j<4; j++ {
			k++
			a[i][j] = k
			fmt.Printf("a[%d][%d]=%d ", i, j, a[i][j])
		}
		fmt.Println()
	}
	fmt.Println("a=",a)

	b := [3][4]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}
	fmt.Println("b=",b)

	//部分初始化
	c := [3][4]int {{1,2,3}, {5,6}, {4}}
	fmt.Println("c=",c)

	e := [3][4]int {1:{1,2,3,4}}
	fmt.Println("e=",e)
}
