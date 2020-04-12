package main

import (
	"regexp"
	"fmt"
)

//const test = "my email is 149015763@qq.com"
const test = `
my email is 149015763@qq.com
email is abcd@qq.com
email is ccc@qq.org.com
`
func main() {
	//res := regexp.MustCompile("149015763@qq.com")
	//res := regexp.MustCompile("[A-Za-z0-9]+@.+\\.com")
	//res := regexp.MustCompile(`[A-Za-z0-9]+@.+\.[A-Za-z0-9]+`)
	res := regexp.MustCompile(`([A-Za-z0-9]+)@(.+)\.([A-Za-z0-9]+)`)
	//match := res.FindString(test)
	//match := res.FindAllString(test, -1)
	match := res.FindAllStringSubmatch(test, -1)
	fmt.Println(match)

	for _, m := range match {
		fmt.Println(m[1])
	}
}
