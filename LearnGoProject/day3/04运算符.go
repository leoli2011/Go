package main

import "fmt"

func main() {
	fmt.Println("1+2 结果：", 1+2)

	A := 20
	A++
	//++A
	fmt.Println("A=", A)
	fmt.Println("5>3 的结果", 5 > 3)

	fmt.Println("!(4>3) 的结果", !(4>3))
}
