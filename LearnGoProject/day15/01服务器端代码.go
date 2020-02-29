package main

import (
	"net"
	"fmt"
)

/*服务器端代码*/
func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8888")
	if err!=nil {
		fmt.Println("failed to listener", err)
		return
	}
	defer ln.Close()

	//阻塞等待用户连接
	conn, err := ln.Accept()
	if err!=nil {
		fmt.Println("failed to accept the connection", err)
		return
	}
	defer conn.Close()

	buff := make([]byte, 1024)
	n, err := conn.Read(buff)
	if err != nil {
		fmt.Println("failed to read the contents", err)
		return
	}
	fmt.Println("n=", n)
	fmt.Println("buff=", string(buff))
}
