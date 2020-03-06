package main

import (
	"net"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9008")
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}

	defer conn.Close()


	requestbuf := "GET /go HTTP/1.1\r\nAccept: image/gif, image/jpeg, image/pjpeg,application/x-ms-application, application/xaml+xml, application/x-ms-xbap,*/*\r\nAccept-Language: zh-Hans-CN,zh-Hans;q=0.8,en-US;q=0.5,en;q=0.3\r\nUser-Agent:Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C;.NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)\r\nAcceptEncoding: gzip, deflate\r\nHost: 127.0.0.1:8000\r\nConnection: Keep-Alive\r\n\r\n"
	n, err := conn.Write([]byte(requestbuf))
	if err != nil {
		fmt.Println("write err", err)
		return
	}
	fmt.Println("send", n, "byte")
	buf := make([]byte, 1024)
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("Read err:", err)
		return
	}
	fmt.Println("get", n, "bytes")
	fmt.Printf("read buf=#%v#\n", string(buf[:n]))
}
