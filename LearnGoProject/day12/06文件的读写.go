package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
)

//文件的写入
func Write_File(path string)  {
	//打开文件，新建文件
	f, err := os.Create(path)
	if (err != nil) {
		fmt.Println("创建文件失败 err=", err)
		return
	}
	//使用完毕 关闭文件
	defer f.Close()

	var buf string
	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("i=%d\n", i)
		n, err := f.WriteString(buf)
		if (err != nil) {
			fmt.Println("write error ", err)
			return
		}
		fmt.Printf("n=%d\n", n)
	}
}

func ReadFile(path string) {
	f, err := os.Open(path)
	if (err != nil) {
		fmt.Println("err=", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 1024*2)
	n, err := f.Read(buf)
	if (err!=nil && err != io.EOF) {
		fmt.Println("err=", err)
		return
	}
	fmt.Printf("buf=%s\n", string(buf[:n]))
}

func ReadFileLine(path string) {
	f, err := os.Open(path)
	if (err != nil) {
		fmt.Println("err=", err)
		return
	}
	defer f.Close()

	//新建一个缓冲区
	r := bufio.NewReader(f)
	for {
		//遇到\n结束读取， 但是\n也进入缓冲区
		buf, err := r.ReadBytes('\n')
		if err!= nil {
			if (err == io.EOF) { //文件已经结束
				break
			}
			fmt.Println("err=", err)
		}
		fmt.Printf("buf=#%s", string(buf))
	}
}
func main() {
	Write_File("./demo.txt")
	fmt.Println("---------------")
	ReadFile("./demo.txt")
	fmt.Println("---------------")
	ReadFileLine("./demo.txt")
}
