package main

import "fmt"

func funcc(c int) {
	fmt.Printf("c=%d\n", c)
}

func funccb(b int) {
	funcc(b-1)
	fmt.Printf("b=%d\n",b)
}
func funcca(a int) {
	funccb(a-1)
	fmt.Printf("a=%d\n", a)
}
func main() {
	funcca(3)
	fmt.Println("main")
}
