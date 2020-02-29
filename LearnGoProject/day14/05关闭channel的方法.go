package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i:=0; i<3; i++ {
			ch <-i
			fmt.Printf("子协程 i=%d\n",i)
		}
		close(ch)
		//ch<-666  //关闭后不能再向channel写数据
	}()

	for {
		//如果ok为true说明管道没有关闭
		if num, ok := <-ch; ok == true {
			fmt.Println("num=", num)
		} else { //管道关闭
			fmt.Println("no more data, break")
			break
		}
	}
}
