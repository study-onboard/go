package ch09

import (
	"fmt"
	"time"
)

func ChannelCommunication() {
	channel := make(chan int)

	for i := 0; i < 10; i++ {
		go goroutine(i, channel)
	}

	time.Sleep(time.Second)

	channel <- 0
	for {
		count := <-channel
		if count > 10 {
			break
		}
		channel <- count
	}
}

func goroutine(number int, channel chan int) {
	fmt.Printf("%d: start.\n", number)

	for {
		count := <-channel
		if count > 10 {
			break
		} else {
			fmt.Printf("%d: readed count - %d\n", number, count)
			count += 1
			fmt.Printf("%d: changed count to - %d\n", number, count)
		}
		channel <- count
	}

	fmt.Printf("%d: complete.\n", number)
}
