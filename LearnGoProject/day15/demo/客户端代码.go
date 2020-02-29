package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer conn.Close()

	go func() {
		buf := make([]byte, 1024)
		//需要在终端执行， 集成环境执行会有问题
		for {
			n1, err := os.Stdin.Read(buf)
			if err != nil {
				fmt.Println("os.stdin err=", err)
				return
			}
			fmt.Println("os read", n1, "bytes")
			n, err := conn.Write(buf[:n1])
			if err != nil {
				fmt.Println("write err=", err)
				return
			}
			fmt.Println("write", n, "bytes")
		}
	}()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err=", err)
			return
		}
		fmt.Printf("got %d bytes  buf=%s", n, buf)
	}
}