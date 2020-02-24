package main

import "fmt"

type Humaner01 interface {
	SayHi()
}
type Personer interface {
	Humaner01
	sing(lrc string)
}

type Student10 struct {
	name string
	id int
}

func (tmp Student10) SayHi()  {
	fmt.Printf("Student10 %s, %d say hi!\n", tmp.name, tmp.id)
}

func (tmp Student10) sing(lrc string)  {
	fmt.Printf("歌词是%s\n", lrc)
}

func main() {
	s := Student10{"mike", 18}
	var i Personer
	i = &s
	i.SayHi()
	i.sing("人民解放歌")

	var h Humaner01
	h = &s
	h.SayHi()

	h = i //子接口可以赋值给父接口
	h.SayHi()
	//i = h //父接口不可以赋值给子接口

}
