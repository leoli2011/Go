package main

import "fmt"

func main() {
	var p *int
	p = nil
	fmt.Println("p=", p)

	//*p = 1000
	var a int
	p = &a
	*p = 1000
	fmt.Println("a=",a)
}
