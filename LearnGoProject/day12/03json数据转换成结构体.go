package main

import (
	"encoding/json"
	"fmt"
)

type IT1 struct {
	Company string
	Isok 	bool
	Price   float64
	Subjects []string
}

type IT2 struct {
	Company string
}

func main()  {
	jsonbuf := `{"company":"wg","isok":true,"price":99,"subjects":["go","docker","fabric","Test"]}`
	var iter IT1
	err := json.Unmarshal([]byte(jsonbuf), &iter)
	if (err != nil) {
		fmt.Println(err)
		return
	}
	//fmt.Println("iter=", iter)
	fmt.Printf("struct=%+v\n", iter)

	var tmp IT2
	err = json.Unmarshal([]byte(jsonbuf), &tmp) //结构体中有部分匹配也可以
	if (err != nil) {
		fmt.Println(err)
		return
	}
	//fmt.Println("iter=", iter)
	fmt.Printf("struct=%+v\n", tmp)

}