package main

import (
	"fmt"
	"errors"
)

func main() {
	//var err1 error = fmt.Errorf("%s", "this is a normal error")
	//自动推导方式
	err1 := fmt.Errorf("%s", "this is a normal error")
	fmt.Println("err1=", err1)

	err2 := errors.New("this is a abnormal error")
	fmt.Println("err2=", err2)

}
