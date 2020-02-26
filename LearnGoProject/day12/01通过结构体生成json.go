package main

import (
	"encoding/json"
	"fmt"
)

/*type IT struct {
	Company string
	Subjects []string
	Isok 	bool
	Price   float64
}*/

/*
type IT struct {
	Company string `json:"company"` // 二次编码，Company小写
	Subjects []string `json:"subjects"`
	Isok 	bool `json:"isok"`
	Price   float64 `json:"price,string"`
}
*/

type IT struct {
	Company string `json:"-"` //-此字段不会输入到屏幕上边
	Subjects []string `json:"subjects"`
	Isok 	bool `json:"isok"`
	Price   float64 `json:"price,string"`
}

func main()  {
	s := IT{"topsec", []string{"go", "docker", "fabric", "Test"}, true, 99}
	//内置方法，不需要依赖外部库
	//buf,err := json.Marshal(s)
	buf,err := json.MarshalIndent(s, "**", "	") //json的格式编码
	if err==nil {
		fmt.Printf("buf=%s", buf)
	} else {
		fmt.Println("err=%s", err)
	}
}
