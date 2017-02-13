package Buffer

import (
	"mine/mineManager"
	"mine/waterCtrl"
	"mine/methaneCtrl"
)

func main() {

	methaneChan := make(chan int)
	waterChan   := make(chan int)
	abort       := make(chan int)


	go mineManager.Manager(abort, methaneChan, waterChan)
	go waterCtrl.Control(waterChan)
	go methaneCtrl.Control(methaneChan)

	<-abort

}


