package main

import "fmt"

func MyFunc14(args ...int) { //不定参数的整形 ...type 不定参数类型
	fmt.Println("len(args)=", len(args))
	for i:=0; i<len(args); i++ {
		fmt.Printf("args[%d]=%d\n",i, args[i])
	}
	fmt.Println("=====================")
	for i,da := range args {//第一个代表下标，第二个代表下标对应的值
		fmt.Printf("args[%d]=%d\n", i, da)
	}
}

func Myfunc15(a int, args ...int)  { //不定参数只能放在形参中的最后一个
	fmt.Println("len(args)=", len(args))
}
func main() {
//	var a, b, c int
//	MyFunc14(1,2,7)
//	MyFunc14(a, b, c)
	MyFunc14();
	Myfunc15(1, 2)
	Myfunc15(2)

}
