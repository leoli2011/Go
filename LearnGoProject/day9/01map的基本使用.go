package main

import "fmt"

func main() {
	//定义一个map变量，类型为map[int]string
	//引用类型不能作为map的key
	//使用map一定要初始化给分配空间
	var m1 map[int]string = make(map[int]string)
	fmt.Println("m1=", m1)

	m1[1] = "jake"
	fmt.Println("m1=", m1)

	m2 := make(map[int]string) //make方式创建然后自动推导类型
	fmt.Println("m2len=", len(m2))
	m2[1] = "mike"
	fmt.Println("m2=", m2)
	fmt.Println("m2len=", len(m2))

	fmt.Println("-----------------")
	//map 先给map指定一个可以容纳长度，一旦超过这个长度 重新分配底层空间
	//map的空间分配和切片相似
	m3 := make(map[int]string, 2)
	m3[1] = "mile"
	m3[2] = "jack"
	m3[3] = "go"
	fmt.Println("m3=", m3)
	fmt.Println("len=", len(m3))
	fmt.Println("-----------------")

	//map可以直接初始化在声明的时候
	m4 := map[int]string{1:"mile", 2:"jason", 3:"ruby"}
	fmt.Println("m4=", m4)

}
