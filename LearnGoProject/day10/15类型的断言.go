package main

import "fmt"

type Student11 struct {
	name string
	id int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "abc"
	i[2] = Student11{"mike", 777}

	//类型的查询， 类型的断言
	//index是返回下标  data对应的值为i[0], i[1], i[2]
	for index, data := range i {
		//第一个返回是值， ok返回是真还是假
		if value, ok := data.(int); ok == true {
			fmt.Printf("i[%d]的类型为int, 值为%d\n", index,value)
		} else if value, ok := data.(string); ok == true{
			fmt.Printf("i[%d]的类型为string, 值为%s\n", index,value)
		} else if value, ok := data.(Student11);ok == true {
			fmt.Printf("i[%d]的类型为student, 内容为" +
				"name=%s, id=%d\n", index,value.name, value.id)
		}
	}
	fmt.Println("-------------第二种判断空接口类型的方式----")
	for index, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Printf("i[%d]的类型为int, 值为%d\n", index,value)
		case string:
			fmt.Printf("i[%d]的类型为string, 值为%s\n", index,value)
		case Student11:
			fmt.Printf("i[%d]的类型为student, 内容为" +
				"name=%s, id=%d\n", index,value.name, value.id)
		}
	}
}