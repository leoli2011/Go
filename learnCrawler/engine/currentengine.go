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
	ConfigureMasterWorkChannel(chan Request)
}
func (e Concurrentengine) Run(seed... Request)  {
	in := make(chan Request)
	out := make(chan ParseResult)

	e.Sche.ConfigureMasterWorkChannel(in)

	//启动worker
	for i:=0; i<e.Workercnt; i++ {
		createWorker(in, out)
	}

	//把request放入channel里边
	for _, r := range seed {
		e.Sche.Submit(r)
	}

	//循环遍历结果
	for {
		result := <-out

		for _, item := range result.Items {
			log.Printf("Got item:%v\n", item)
		}

		for _, request := range result.Requests {
			e.Sche.Submit(request)
		}
	}
}

/*
创建工作任务
*/
func createWorker(in chan Request, out chan ParseResult)  {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}