package main

import (
	"fmt"
)

type Worker struct {
	name string
	sex byte
	age int
}

func (a Worker) Printinfo() {
	fmt.Println("Worker=", a)
}

func Printinfo1(a Worker) {
	fmt.Println("OOP: ", a)
}

/*func (p Worker) SetInfo(name string, sex byte, age int)  {
	p.name = name
	p.sex = sex
	p.age = age
	fmt.Println("SetInfo: p=", p)
}*/

func (p *Worker) SetInfo(name string, sex byte, age int)  {
	p.name = name
	p.sex = sex
	p.age = age
	fmt.Println("SetInfo: p=", p)
}

func main() {
	w := Worker{"zhangsan", 'M', 20}

	//面向对象
	w.Printinfo()
	//面向过程
	Printinfo1(w);
	fmt.Println("=========================")
	var p Worker
	p.SetInfo("Tom", 'F', 30)
	p.Printinfo()
	//(&p).SetInfo("john", 'M', 29)
	p.Printinfo()

}
