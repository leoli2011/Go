package main

import (
	"fmt"
	"os"
	"net"
	"io"
)

func sendFile(fn *os.File, conn net.Conn) {
	buf := make([]byte, 1024 *4)
	for {
		n, err := fn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("finish send file", err)
				break
			}
			fmt.Println("fn.read failed err=", err)
			return
		}
		conn.Write(buf[:n])
	}
	return
}

func main() {
	fmt.Println("请输入传送文件名：")
	var path string
	fmt.Scan(&path)

	fmt.Printf("#%v#\n", path)
	fi, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat err=", err)
		return
	}
	fmt.Println(fi.Name(), fi.Size())

	fh,err := os.Open(path)
	if err != nil {
		fmt.Println("os.Open err=", err)
		return
	}

	defer fh.Close()

	conn, err := net.Dial("tcp", "127.0.0.1:7000")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	n, err := conn.Write([]byte(fi.Name()))
	if err != nil {
		fmt.Printf("failed to send file name, err=", err)
		return
	}
	fmt.Printf("send %d to server\n", n)

	buf := make([]byte, 1024)
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}

	fmt.Println("client got ", n , "bytes")
	fmt.Println("contents= ", string(buf))

	if "OK" == string(buf[:n]) {
		sendFile(fh, conn)
	}

	fmt.Println("finish send file")

}
