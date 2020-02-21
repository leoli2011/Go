package main

import "fmt"

func main() {
	var id [50] int //数组同一类型集合

	for i := 0; i < len(id); i++ {
		id[i] = i+1
		fmt.Printf("id[%d]=%d\n", i, id[i])
	}

}
