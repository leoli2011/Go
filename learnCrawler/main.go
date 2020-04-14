package main

import (
	"learnCrawler/zhenai/parse"
	"learnCrawler/engine"
)

func main() {
	engine.Run(engine.Request{
		"http://www.zhenai.com/zhenghun",
		parse.ParseCity,
	})
}
