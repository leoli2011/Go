package main

import "fmt"

func testsum(a int) int {
	if a==1 {
		return 1
	}
	return a+testsum(a-1)
}

func testsum1(a int) int {
	if a==100 {
		return 100
	}
	return a+testsum1(a+1)
}
func main() {
	var sum int
	sum = testsum(100)
	fmt.Println("sum=", sum)
	sum = testsum1(1)
	fmt.Println("sum1=", sum)
}
