package main

import (
	"net"
	"fmt"
	"io"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:9005")
	if err != nil {
		fmt.Println("lister err:", err)
		return
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err!= nil {
		fmt.Println("accept err:", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err!= nil {
		if err == io.EOF {
			fmt.Println("finish:", err)
			return
		}
		fmt.Println("read err:", err)
		return
	}

	fmt.Println("got", n, "bytes from client")
	fmt.Printf("buf=#%v#", string(buf[:n]))

	}
