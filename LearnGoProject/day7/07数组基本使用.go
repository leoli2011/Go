package main

import "fmt"

func main() {
	var a [10]int
	var b [5]int
	fmt.Printf("len(a)=%d, len(b)=%d\n",  len(a), len(b))
	//定义的数组个数必须是常量

/*	n := 10
	var c [n]int //报错： non-constant array bound n
	fmt.Printf("len(c)=%d\n",  len(c))*/

	//数组元素下标从0开始到len-1
	//变量可以作为数组下标

	for i := 0; i < len(a); i++ {
		a[i] = i+1
	}

	for i, data := range a {
		fmt.Printf("a[%d]=%d\n", i, data)
	}




}
