package waterCtrl

import (
	"math/rand"
	"time"
	"fmt"
)

func Control (waterChan chan<- int) {

	for{
		t := rand.Int63n(10 * int64(time.Second))
		time.Sleep(time.Duration(t))

		sensor := rand.Intn(60)
		fmt.Println(">>>> Water: ", sensor)
		waterChan <- sensor
	}

}