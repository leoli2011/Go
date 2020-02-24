package main

import "fmt"

type Worker02 struct {
	name string
	sex byte
	age int
}

// 接收者为普通变量，　非指针，　值语义　　一份拷贝
func (w Worker02) SetWorkInfo() {
	fmt.Printf("setworkinfo:  w address =%p\n", &w)
}

//引用传递
func (w *Worker02) SetWorkPoint() {
	fmt.Printf("setworkpoint:  w address =%p\n", w)
}

func main() {
	/*
	w1 := &Worker02{"marry", 'M', 33}
	//结构体变量是一个指针类型的变量，它能够调用方法，这些可以调用的方法简称方法集
	w1.SetWorkPoint()
	(*w1).SetWorkPoint() //会把(*w)转换为w再调用，等价和上边
	//内部会先把指针w1转换成*w再调用接收者为普通的变量的方法
	//(*w1).SetWorkInfo()
	w1.SetWorkInfo()
	*/
	w2 := Worker02{"marry", 'M', 33}
	w2.SetWorkPoint() //w 会先转换地址&w在调用方法
	(&w2).SetWorkPoint()
	w2.SetWorkInfo()


}
