package engine

import "log"

type Concurrentengine struct {
	Sche  Scheduler
	Workercnt int
}

type Scheduler interface {
	//传入request方法
	Submit(Request)
	//创建一个channel
	//ConfigureMasterWorkChannel(chan Request)
	//传入线程对应的接收Request 的channel
	WorkerReady(chan Request)
	Run()
}
func (e Concurrentengine) Run(seed... Request)  {
	out := make(chan ParseResult)
	e.Sche.Run()

	//启动worker
	for i:=0; i<e.Workercnt; i++ {
		createWorker(out, e.Sche)
	}

	//把request放入channel里边
	for _, r := range seed {
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
			e.Sche.Submit(request)
		}
	}
}

/*
创建工作任务
*/
func createWorker(out chan ParseResult, s Scheduler)  {
	in := make(chan Request)
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