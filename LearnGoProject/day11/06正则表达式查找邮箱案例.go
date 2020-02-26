package main

import (
	"regexp"
	"fmt"
)

const text = `
my email is 530979104@qq.com
email is adbd@qq.com
email is ccc@qqq.org.com`

func main() {
	//reg := regexp.MustCompile("530979104@qq.com")
	reg := regexp.MustCompile("([0-9a-zA-Z]+)@([0-9a-zA-Z]+)\\.([0-9a-zA-Z.]+)+")
	//result := reg.FindAllString(text, -1)
	result := reg.FindAllStringSubmatch(text, -1)
	fmt.Printf("result:=%s\n", result)

	for _, m := range result {
		fmt.Println("data=", m[1])
	}
}
