package main

import (
	"net"
	"fmt"
)

func main()  {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer conn.Close()

	n, err := conn.Write([]byte("hello server"))
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Printf("send %d byte to server", n)
}
