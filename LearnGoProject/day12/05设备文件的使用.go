package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Stdout.Close()  //关闭后无法输出
	fmt.Println("are you ok?")
	os.Stdout.WriteString("ARE YOU OK?\n")

	// os.Stdin
	var  a int
	fmt.Println("请输入一个A")
	fmt.Scan(&a)
	fmt.Println("a=", a)
}
