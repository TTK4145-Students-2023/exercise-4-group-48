package timer

import (
	"Project/network"
	"fmt"
	"os"
	"time"
)

func A(file *os.File) {
	// Initialize timer
	var duration = 1 * time.Second
	timer1 := time.NewTimer(duration)
	var number = 5
	fmt.Println("start timer")

	defer file.Close()

	for i := 0; i < number; i++ {
		// Count
		timer1.Reset(duration)
		<-timer1.C
		fmt.Println("Counting: ", i)

		// Update file
		text := fmt.Sprintf("%d", i)
		text = text + "\n"
		file.WriteString(text)
		fmt.Println("Number written to file:", i)

		// Send heartbeat
		var heartbeat string = "heartbeat"
		network.UDPSend(heartbeat)
	}

}
