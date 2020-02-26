package main

import (
	"regexp"
	"fmt"
)

func main() {
	reg := regexp.MustCompile(`^z.*l$`)
	result := reg.FindAllString("zhangsanl zhangsanl", -1)
	fmt.Printf("result=%s\n", result)

	reg = regexp.MustCompile(`^z(.*)l$`)
	result = reg.FindAllString("zhangsanl", -1)
	fmt.Printf("result=%s\n", result)
}
