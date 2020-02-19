package main

import "fmt"

func main() {
	str := "abc"
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d]=%c \n", i, str[i])
	}

	fmt.Println("--------------------------------")

	//range迭代打每个元素，默认返回两个值，第一个是元素的位置，第二个是元素本身
	for i,data :=range str {
		fmt.Printf("str[%d]=%c \n", i, data)
	}

	for i :=range str { //这种形式默认第二个返回值丢弃， 返回元素的位置(下标)
		fmt.Printf("str[%d]=%c \n", i, str[i])
	}

	fmt.Println("\n6666666666666666666666666")
	for i,_ :=range str { //用_把第二个返回值丢弃
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}

}
