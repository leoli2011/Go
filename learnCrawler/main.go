package main

import (

	"learnCrawler/fetch"
	//"fmt"
	"learnCrawler/zhenai/parse"
)

func main() {
	all, err := fetch.Fetch("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	parse.ParseCity(all)
	//fmt.Printf("###%s###\n", all)
}
