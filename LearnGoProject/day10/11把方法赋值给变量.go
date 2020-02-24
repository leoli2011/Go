package main

import "fmt"

type Person7 struct {
	name string
	sex byte
	age int
}

func (p Person7) SetInfoValue() {
	fmt.Printf("SetInfoValue:%p, %v\n", &p, p)
}

func (p *Person7) SetInfoPoint()  {
	fmt.Printf("SetInfoPoint:%p, %v\n", p, p)
}

func main() {
	p := Person7{"mike", 'm', 18}
	p.SetInfoPoint()

	pFunc := p.SetInfoPoint//这个是方法值，调用函数时，无需传递接收者
	pFunc()

	PSetInfo := p.SetInfoValue
	PSetInfo() //这个就等价于p.SetInfoValue()
}
