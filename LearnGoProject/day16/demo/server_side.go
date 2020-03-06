package main

import (
	"net"
	"fmt"
	"os"
	"io"
)

func rece_File(filename string, conn net.Conn) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("failed to create file err=", err)
		return
	}

	buf := make([]byte, 1024 *4)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read err", err)
			return
		}
		fmt.Println("got ", n, "bytes")
		f.Write(buf[:n])
	}
}
func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:7000")
	if err != nil {
		fmt.Println("listener err=", err)
		return
	}
	defer ln.Close()

	conn,err := ln.Accept()
	if err != nil {
		fmt.Println("accept err=", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	n,err := conn.Read(buf)
	if err != nil {
		fmt.Println("read err", err)
		return
	}
	fmt.Println("we got", n, "bytes")
	fmt.Println(string(buf))

	conn.Write([]byte("OK"))
	rece_File(string(buf[:n]), conn)

}
