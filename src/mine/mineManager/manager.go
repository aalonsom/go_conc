package mineManager

import "fmt"

func motorOn () {
	fmt.Println("!!!! Motor on")

}

func motorOff() {
	fmt.Println("!!!! Motor off")
}

func Manager (abort chan<- int, methaneChan <-chan int, waterChan <-chan int) {

	// abort: for stopping the main. Not used

	const methaneThreshold   = 30
	const waterLowThreshold  = 20
	const waterHighThreshold = 40

	var methaneLevel  = 0
	var waterLevel    = 0
	var motorWorking  = false

	// wait the values from sensors
	for {
		select {
		case methaneLevel = <-methaneChan:

		case waterLevel   = <-waterChan:
		}

		switch motorWorking {

		case false:
			if waterLevel > waterHighThreshold && methaneLevel < methaneThreshold {
				motorWorking = true
				motorOn()
			}
		case true:
			if (waterLevel < waterLowThreshold || methaneThreshold > methaneThreshold) {
				motorWorking = false
				motorOff()
			}
		}
	}
}
