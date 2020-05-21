package scheduler

import "learnCrawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

/**
给workerchan赋值
*/

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {s.workerChan <- r}()
//	s.workerChan <- r
}

/**
给workerChan初始化
*/
func (s *SimpleScheduler) ConfigureMasterWorkChannel(c chan engine.Request) {
	s.workerChan = c
}


