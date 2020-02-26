package main

import (
	"regexp"
	"fmt"
)

func main() {
	buf := "43.14 567 adfdf  1.27 7 8.9 ddddsss 6.66 0.78"
	reg := regexp.MustCompile(`\d+\.\d+`)
	//result := reg.FindAllString(buf, -1)
	result := reg.FindAllStringSubmatch(buf, -1) //把每一个匹配的单独放到一个数组方便以后处理
	fmt.Printf("result=%s", result)
}
