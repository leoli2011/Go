package main

import (
	"net"
	"fmt"
	"strings"
)

func handleConnection(conn net.Conn )  {
	defer conn.Close()//函数调用完毕，自动关闭conn
	//获取客户的网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println("client addr =", addr)
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err=", err)
		}
		fmt.Println("got", n, "bytes")
		fmt.Printf("string:%s", string(buf))
		if "exit" == string(buf[0:n-1-1]) {
			fmt.Println("exit")
			break
		}

		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer ln.Close()
	//接受多个用户
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("err=", err)
			return
		}
		defer conn.Close()
		//处理用户请求，新建一个协程
		go handleConnection(conn)
	}
}
