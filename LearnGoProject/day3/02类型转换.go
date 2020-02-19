package main

import "fmt"

func main() {
	var ch byte;
	ch = 'a'
	var T int
	T = int(ch)
	//T = ch
	fmt.Printf("T=%d\n", T)

	var flag bool
	flag = true
	fmt.Printf("%t\n", flag)
	//fmt.Printf("%d\n", int(flag))

}