package methaneCtrl

import "time"
import (
	"math/rand"
	"fmt"
)

func Control (methaneChan chan<- int) {

	for{
		t := rand.Int63n(10 * int64(time.Second))
		time.Sleep(time.Duration(t))

		sensor := rand.Intn(60)
		fmt.Println(">>>> Methane: ", sensor)
		methaneChan <- sensor
	}


}
