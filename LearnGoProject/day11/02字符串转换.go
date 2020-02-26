package main

import (
	"strconv"
	"fmt"
)

func main()  {
	//转换成字符后追加到字节数组中
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	slice = strconv.AppendInt(slice, 1234, 10)
	slice = strconv.AppendQuote(slice, "abcdefg")
	fmt.Println("slice=", string(slice))

	//把其他类型转换成字符串
	var str string
	str = strconv.FormatBool(true)
	fmt.Println("str =", str)
	str = strconv.FormatInt(10, 8)
	fmt.Println("str =", str)

	str = strconv.FormatFloat(3.14, 'f', -1, 64)
	fmt.Println("str=", str)
	//常用， 把整形转换为字符串
	str = strconv.Itoa(666)
	fmt.Println("str=", str)

	//把字符串转换成其它类型
	var flag bool
	var err error
	flag, err  = strconv.ParseBool("true")
	if (err == nil) {
		fmt.Println("flag = ", flag)
	} else {
		fmt.Println("err=", err)
	}
	//把字符串转换为整形
	a, err := strconv.Atoi("888")
	if (err == nil) {
		fmt.Println("a = ", a)
	} else {
		fmt.Println("err=", err)
	}
}
