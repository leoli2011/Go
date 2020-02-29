package main

func main() {
	//创建一个管道
	ch := make(chan int)
	//双向channel能隐式转换为单向管道
	var writeCh chan<- int = ch //只能写，不能读
	var readCh <-chan int = ch  //只能读，不能写
	writeCh <- 666
	//<-writeCh
	<-readCh //读
	//readCh <- 666
	//单向channel不能转转换为双向的
	//var ch2 int = writeCh
}
