package main

import (
	"learnCrawler/zhenai/parse"
	"learnCrawler/engine"
	"learnCrawler/scheduler"
)

func main() {
	/*
	engine.Run(engine.Request{
		"http://www.zhenai.com/zhenghun",
		parse.ParseCity,
	})
	*/

	e :=engine.Concurrentengine{
		Sche:&scheduler.SimpleScheduler{},
		Workercnt:2,
	}

	e.Run(engine.Request{
		"http://www.zhenai.com/zhenghun",
		parse.ParseCity,
	})
}
