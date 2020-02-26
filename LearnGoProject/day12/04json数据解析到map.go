package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonbuf := `{"company":"wg","isok":true,"price":99,"subjects":["go","docker","fabric","Test"]}`
	m1 := make(map[string]interface{}, 4)

	err := json.Unmarshal([]byte(jsonbuf), &m1)//传递地址
	if (err != nil) {
		fmt.Println("err=", err)
		return
	}
	fmt.Printf("buf=%+v\n", m1)

	//var str string
	//str = m1["company"]

	for key, data := range m1 {
		switch data.(type) 	{
		case string:
			fmt.Printf("m[%s]的类型为string %s\n", key, data)
		case bool:
			fmt.Printf("m[%s]的类型为bool %v\n", key, data)

		case float64:
			fmt.Printf("m[%s]的类型为float64 %v\n", key, data)

		case interface{}:
			fmt.Printf("m[%s]的类型为interface{} %v\n", key, data)
		case []interface{}:
			fmt.Printf("m[%s]的类型为[]interface{} %v\n", key, data)
		}
	}
}
