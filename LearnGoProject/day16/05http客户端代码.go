package main

import (
	"net/http"
	"fmt"
	"io"
)

func main() {
	//resp, err := http.Get("http://www.baidu.com")
	resp, err := http.Get("http://127.0.0.1:8000")
	if err != nil {
		fmt.Println("http Get err:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("resp.status=", resp.Status)
	fmt.Println("resp.code=", resp.StatusCode)
	fmt.Println("header=", resp.Header)
	//fmt.Println("body=", resp.Body) //body= &{0xc042050300 <nil> <nil>}

	buf := make([]byte, 1024 * 4)
	var tmp string
	for {
		n, err := resp.Body.Read(buf)
		fmt.Println("n=",n)
		if err != nil || n==0 {
			if err == io.EOF{
				tmp += string(buf[:n])
				break
			}
			fmt.Println("body read err", err)
			break
		}
		tmp += string(buf[:n])
	}
	fmt.Println("temp = ", tmp)
}
