package main

import (
//	"bufferMulti/buffer"
	"bufferMulti/producer"
	"bufferMulti/consumer"
)

func main() {


	const nConsumers = 2
	const nProducers = 5

	i := 0
	var abort = make(chan int)

//	go buffer.Buffer(abort)
	for i = 1; i <= nConsumers; i++  {
		go consumer.Consumer(i)
	}
	for i = 1; i <= nProducers; i++ {
		go producer.Producer(i)
	}
	<-abort

}


