package main

import "fmt"

func Maxandmin(a, b int) (max, min int) {
	if (a>b) {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	return
}

func main() {
	var a,b int
	a = 10
	b = 50
	max, min := Maxandmin(a, b)
	fmt.Printf("max=%d, min=%d\n", max, min)
	//通过匿名变量丢弃某个返回值_
	big, _ :=Maxandmin(a, b)
	fmt.Printf("big=%d\n", big)

}
