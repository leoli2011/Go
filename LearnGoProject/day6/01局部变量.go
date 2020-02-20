package main

import (
	"fmt"
)

func main() {
	a := 111
	fmt.Printf("a=%d",a)
	{
		b :=222
		fmt.Println("b=", b)
	}
	//fmt.Println("b=", b)

	if flag :=3; flag==3 {
		fmt.Println("flag=", flag)
	}
}
