package main

import (
	"os"
	"fmt"
)

func main() {
	list := os.Args
	fmt.Printf("len:(%d)\n", len(list))

	if len(list) != 2 {
		fmt.Println("xxxx usage")
		return
	}

	name := list[1]
	info, err := os.Stat(name)
	if err != nil {
		fmt.Println("err=",err)
		return
	}
	fmt.Println(info.Name(), info.Size())
}
