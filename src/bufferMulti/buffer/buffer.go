package buffer

import "fmt"

//import "sync"

const maxB = 5

var buffer int  = 0
var exist  bool = false
var nDatos int  = 0

//var mutex sync.Mutex
var bufferCh = make(chan int, maxB)
var consCh = make(chan int)

func Put(data int) {
	bufferCh <- data
	fmt.Println("NBUFFER (Put):", len(bufferCh))
}

func Get() (data int) {
	data = <- bufferCh
	fmt.Println("NBUFFER (Get):", len(bufferCh))
	return data
}
