package main

import (
	"fmt"
)

func main() {
	m1 := map[int]string {1:"mike", 2:"jack", 3:"tom"}

	//第一个返回值为key   第二个返回值为value
	//遍历结果是无序的
	for key, value := range m1 {
		fmt.Printf("%d===========>%s\n", key, value)
	}

	fmt.Println("----------------------")
	//m1[0] = "python"
	//判断一个key是否存在的方法
	//value, ok := m1[0]
	if value, ok := m1[0]; ok == true {
		fmt.Printf("m1[0]=%s", value)
	} else {
		fmt.Printf("key不存在\n")
	}

	// map的删除
	delete(m1, 2)  //删除key为２的内容
	fmt.Println("m1=", m1)
}
