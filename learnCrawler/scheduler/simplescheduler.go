package scheduler

import "learnCrawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

/**
给workerchan赋值
*/

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {s.workerChan <- r}()
//	s.workerChan <- r
}


