package main

import (
	"os"
	"fmt"
)

func main() {
	list := os.Args

	n := len(list)

	fmt.Printf("len=%d\n", n)

	for i := 0; i < n; i++ {
		fmt.Printf("list[%d]=%s \n", i, list[i])
	}
	fmt.Println("------------------------")

	for i,data := range list {
		fmt.Printf("i=%d, data=%s\n",i, data)
	}
}
