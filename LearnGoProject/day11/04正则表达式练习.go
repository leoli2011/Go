package main

import (
	"regexp"
	"fmt"
)

func main() {
	buf := "abc azc a7c aac 888 a9c tac"
	//1， 编译匹配规则
	reg := regexp.MustCompile("a(.)c") //把满足条件的放在一组
	//2, 拿规则匹配所有字符串
	//result := reg.FindAllString(buf, 2)
	result := reg.FindAllStringSubmatch(buf, -1)
	fmt.Printf("result=%s\n", result[0][1])

	//reg2 := regexp.MustCompile("a([0-9]+)c")
	reg2 := regexp.MustCompile(`a(\d+)c`)
	result2 := reg2.FindAllStringSubmatch(buf, -1)
	fmt.Printf("result=%s\n", result2)

	//reg3 := regexp.MustCompile(`a[a-zA-Z]c`)
	reg3 := regexp.MustCompile(`a\wc`)
	result3 := reg3.FindAllStringSubmatch(buf, -1)
	fmt.Printf("result=%s\n", result3)

}
