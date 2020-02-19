package main

import "fmt"

func main()  {
	var b int
	b = 3
	switch b {//在case里默认加了break语句
	 //加fallthrough继续执行case后的语句
	case 1:
		fmt.Println("这个数字是1")
	case 2:
		fmt.Println("这个数字是2")
	case 3:
		fmt.Println("这个数字是3")
		fallthrough
	case 4:
		fmt.Println("这个数字是4")
		fallthrough//只要不跳出switch语句，后面无条件执行
	default:
		fmt.Println("这个数字是", b)
	}
}
