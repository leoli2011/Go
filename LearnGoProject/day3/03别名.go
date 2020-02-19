package main

import "fmt"

func main() {
	type bigint int64
	//var a int64
	var a bigint
	fmt.Printf("a type is %T\n", a)
	type (
		long int64
		char byte
	)
	var ch char='b'
	var b long = 11
	fmt.Printf("b=%d, ch=%c", b, ch)
}
