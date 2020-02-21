package main

import "fmt"

func main() {
	s := make([]int, 0, 1)  //类型 长度  容量
	oldcap := cap(s)
	for i := 0; i < 20; i++ {
		s = append(s, i)
		if newcap:=cap(s); oldcap < newcap {
			fmt.Printf("oldcap:%d-------->%d newcap\n", oldcap, newcap)
			oldcap = newcap
		}
	}
}