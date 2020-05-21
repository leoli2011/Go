package main

import (
	"learnCrawler/zhenai/parse"
	"learnCrawler/engine"
	"learnCrawler/scheduler"
	"learnCrawler/persist"
)

func main() {
	/*
	engine.Run(engine.Request{
		"http://www.zhenai.com/zhenghun",
		parse.ParseCity,
	})
	*/

	e :=engine.Concurrentengine{
		Sche:&scheduler.QueuedScheduler{},
		Workercnt:2,
		Itemchan: persist.ItemSave(),
	}

	e.Run(engine.Request{
		"http://www.zhenai.com/zhenghun/aba/",
		parse.ParsecityUser,
	})
}
