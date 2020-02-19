package main

import "fmt"

func main() {
	switch num:=4; num {
	case 1:
		fmt.Println("1的数字")
	case 2:
		fmt.Println("2的数字")
	case 3,4,5: //相同的语句可以合并
		fmt.Println("zzz的数字")
	case 6:
		fmt.Println("1的数字")
	default:
		fmt.Println("XXX的数字")
	}
	fmt.Println("-------------------")
	num:=4
	switch { //可以没有条件
	case num>60: //case后边可以跟上条件语句
		fmt.Println("大于60的数字")
	case num>90:
		fmt.Println("大于90的数字")
	case num==100: //相同的语句可以合并
		fmt.Println("100的数字")
	default:
		fmt.Println("其它的数字")
	}
}
