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
	}()

	for num := range ch{
		fmt.Println("num=", num)
	}

}
