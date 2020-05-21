package persist

import "log"

func ItemSave() chan interface{}  {
	out := make(chan interface{})

	go func() {
		cnt := 0
		for{
			item := <- out
			log.Printf("Item Saver: go item #%d:%v",  cnt, item)
			cnt++
		}

	}()
	return out
}
