package buffer

//import "sync"

//const maxB = 5

var buffer int  = 0
var exist  bool = false

//var mutex sync.Mutex
var prodCh = make(chan int)
var consCh = make(chan int)

func Put(data int) {
	prodCh<- data
}

func Get() (data int) {
	data = <-consCh
	return data
}

func Buffer(abort chan<- int) {
	for {
		if exist {
			// consumer
			consCh <- buffer
			exist = false
		} else {
			// producer
			buffer = <-prodCh
			exist = true
		}
	}
}



/*func put(data int) {
	buffer = <-consCh
}

func get() {//(data int){
	prodCh <- buffer
}

func Init (pCh, cCh) {
	prodCh = pCh
	consCh = cCh
*/



