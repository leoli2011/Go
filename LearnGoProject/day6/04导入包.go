package main

/*import (
	_ "fmt" //使用下划线不使用也不会报错 相当于忽略此包
)*/

import . "fmt" //无需通过包名调用，不建议这样操作

import (
	io "fmt" //给包起个别名叫IO
)


func main() {
	io.Println("io alias fmt")
	Println("no package name")
}
