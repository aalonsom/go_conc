package consumer

import (
	"time"
	"fmt"
	"math/rand"
	"bufferSingle/buffer"
)

func Consumer (id int) {

//	pos := 0
	data := 0

	for{
		t := rand.Int63n(10 * int64(time.Second))
		time.Sleep(time.Duration(t))

		data = buffer.Get()
		fmt.Println("<<<<< Consumer. Id: ", id, "Data: " , data)
	}

}
