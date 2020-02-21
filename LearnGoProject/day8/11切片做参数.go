package main

import (
	"math/rand"
	"time"
	"fmt"
)

func initdata(s []int) {
	rand.Seed(time.Now().UnixNano())
	for i,_ := range s {
		s[i] = rand.Intn(100)
	}
}

func bubblesort(s []int) {
	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := 0; j<n-1-i; j++  {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1],s[j]
			}
		}
	}
}

func main() {
	n := 10
	s1 := make([]int, n)
	initdata(s1)
	fmt.Println("排序前:-------------")
	fmt.Println(s1)
	bubblesort(s1)
	fmt.Println("排序后：")
	fmt.Println(s1)
}
