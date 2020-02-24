package main

import "fmt"

type Worker01 struct {
	name string
	sex byte
	age int
}

// 接收者为普通变量，　非指针，　值语义　　一份拷贝
func (w Worker01) SetWorkInfo(name string, sex byte, a int) {
	w.name = name
	w.sex = sex
	w.age = a
	fmt.Printf("setworkinfo:  w address =%p\n", &w)
}

//引用传递
func (w *Worker01) SetWorkPoint(name string, sex byte, a int) {
	w.name = name
	w.sex = sex
	w.age = a
	fmt.Printf("setworkinfo:  w address =%p\n", w)
}

func main() {
	w1 := Worker01{"mike",'M', 18}
	fmt.Printf("main:  w address =%p\n", &w1)
	//w1.SetWorkInfo("marry", 'F', 29)
	//fmt.Printf("main:  w address =%p\n", &w1)
	//(&w1).SetWorkPoint("marry", 'F', 29)
	w1.SetWorkPoint("marry", 'F', 29)

	fmt.Printf("&w = %p\n", &w1)
	fmt.Println("w1=", w1)

}
