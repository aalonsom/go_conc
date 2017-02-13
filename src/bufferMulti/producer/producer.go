package producer

import (
	"time"
	"fmt"
	"math/rand"
	"bufferMulti/buffer"
)

func Producer (id int) {

	pos := 0

	for{
		t := rand.Int63n(10 * int64(time.Second))
		time.Sleep(time.Duration(t))

		pos ++
		data := id * 1000 + pos
		fmt.Println(">>>>> Producer. Id: ", id, "Data: " , data)
		buffer.Put(data)
	}

}

