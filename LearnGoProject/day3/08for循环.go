package main

import "fmt"

func main() {
	sum := 0
	//var i int
	for i:=1; i<=100; i++ {
		sum = sum + i
	}

	fmt.Println("sum=", sum)
}
