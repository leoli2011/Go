package main

import "fmt"

func main() {
	var p *int
	p  = new(int)
	fmt.Printf("p=%v, *p=%d\n", p, *p)

	*p = 6000
	fmt.Printf("p=%v, *p=%d\n",p, *p)

	q := new(int)
	*q = 8888
	fmt.Printf("q=%v, *q=%d\n", q, *q)

}
