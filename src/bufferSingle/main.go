package main

import (
	"bufferSingle/buffer"
	"bufferSingle/producer"
	"bufferSingle/consumer"
)

func main() {


	const nConsumers = 5
	const nProducers = 5

	i := 0
	var abort = make(chan int)

	go buffer.Buffer(abort)
	for i = 0; i < nConsumers; i++  {
		go consumer.Consumer(i)
	}
	for i = 0; i < nProducers; i++ {
		go producer.Producer(1)
	}
	<-abort

}


