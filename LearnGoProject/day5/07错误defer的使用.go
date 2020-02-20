package main

import "fmt"

func test01(x int) {
	result := 1000/x
	fmt.Println("result=", result)
}

func main() {
	defer fmt.Println("aaaaaa")
	defer fmt.Println("bbbbbb")
	defer test01(0)
	defer fmt.Println("cccccc")
}