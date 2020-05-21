package engine

import "log"

type Concurrentengine struct {
	Sche  Scheduler
	Workercnt int
}

type Scheduler interface {
	ReadyNotifiler
	//传入request方法
	Submit(Request)

	Run()
	WorkerChan() chan Request
}

type ReadyNotifiler interface {
	//传入线程对应的接收Request 的channel
	WorkerReady(chan Request)
}

func (e Concurrentengine) Run(seed... Request)  {
	out := make(chan ParseResult)
	e.Sche.Run()

	//启动worker
	for i:=0; i<e.Workercnt; i++ {
		createWorker(e.Sche.WorkerChan(), out, e.Sche)
	}

	//把request放入channel里边
	for _, r := range seed {
		if isDuplicated(r.Url) {
			continue
		}
		e.Sche.Submit(r)
	}

	itemCount := 0
	//循环遍历结果
	for {
		result := <-out

		for _, item := range result.Items {
			log.Printf("Got item:#%d, %v\n", itemCount, item)
			itemCount++
		}

		for _, request := range result.Requests {
			if isDuplicated(request.Url) {
				continue
			}
			e.Sche.Submit(request)
		}
	}
}

var visitedUrls = make(map[string]bool)

func isDuplicated(url string) bool  {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true

	return false
}

/*
创建工作任务
*/
func createWorker(in chan Request, out chan ParseResult, s ReadyNotifiler)  {
	go func() {
		for {
			s.WorkerReady(in)
			request := <- in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}