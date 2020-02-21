package main

import (
	"fmt"
)

func main() {
	srcslice := []int{1,2}
	dstslice := []int{7,7,7,7}
	//copy(dstslice,srcslice)
	copy(srcslice,dstslice )
	fmt.Println(srcslice)
}
