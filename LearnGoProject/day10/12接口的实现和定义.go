package main

import "fmt"

//定义一个接口类型
type Humaner interface {
	//方法，　只有声明，没有实现，由别的类型(自定义)类型实现
	sayHi()
}

type Student09 struct {
	name string
	id int
}

func (tmp Student09) sayHi() {
	fmt.Printf("Student09[%s:%d]\n", tmp.name, tmp.id)
}

type Teacher09 struct {
	addr string
	id  int
}

func (tmp Teacher09) sayHi()  {
	fmt.Printf("Teacher09[%s:%d]\n", tmp.addr, tmp.id)
}

type Worker04 struct {
	id int
}

func (tmp Worker04) sayHi() {
	fmt.Printf("Worker04 [%d] sayHi\n", tmp.id)
}

//定义一个普通函数，函数的参数为接口类型
//只有一个函数，　可以表现为不同的形式，这就是多态

func WhoSayHi(i Humaner) {//接口作为参数
	i.sayHi()
}
func main() {
	//一个接口类型变量
	var i Humaner
	//只有实现了此接口的方法的类型，　那么这个类型的变量才能赋值给接口变量
	s := &Student09{"mike", 777}
	i = s
	i.sayHi()
	t := &Teacher09{"bj", 888}
	i = t
	i.sayHi()

	w := Worker04{8888}
	i = w
	w.sayHi()

	fmt.Println("------------")
	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(w)

	fmt.Println("××××××××××××")
	//创建一个切片, 切片放入是接口
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = w
	for _, i := range x {
		i.sayHi()
	}
}