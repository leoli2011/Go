package main

import (
	"strings"
	"fmt"
)

func main() {
	//contains包含返回true不包含返回false
	result := strings.Contains("hellogo", "hello")
	fmt.Println("result=",result)
	result = strings.Contains("hellogo", "abc")
	fmt.Println("result=",result)

	//join  组合
	s := []string{"hello","abc", "mike", "go"}
	buf := strings.Join(s, "_")
	fmt.Println("buf=", buf)

	//index查找子串位置
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "llll"))

	//repeat
	buf = strings.Repeat("go", 5)
	fmt.Println("buf=",buf)

	//split
	buf = "hello_abc_mike_go"
	s2 := strings.Split(buf, "_")
	fmt.Println("s2=", s2)

	//trim
	buf = strings.Trim("   are you ok    ", "  ")
	fmt.Printf("buf=#%s#\n", buf)

	//fields
	slice :=  strings.Fields("   are you ok??? ")
	fmt.Println("buf=", slice)

	for i, data := range slice {
		fmt.Println(i, data)
	}

}
