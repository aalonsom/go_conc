package buffer

import "sync"

const maxB = 5

var mutex sync.Mutex
var prodCh chan make(chan int, maxB)
var consCh chan make()

func Put(data int) {


}

func Get() (int) {

}

func Init (pCh, cCh) {
	prodCh = pCh
	consCh = cCh


}

