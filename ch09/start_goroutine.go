package ch09

import (
	"fmt"
	"time"
)

func StartGoroutine() {
	go running("Masaga")
	go running("Mona")
	go running("Bowen")

	go func() {
		for {
			fmt.Println("-------------------------------------------")
			time.Sleep(time.Second * 3)
		}
	}()

	var input string
	fmt.Scanln(&input)
	fmt.Printf("input: %s\n", input)
}

func running(name string) {
	var times int
	for {
		times++
		fmt.Printf("%s: %d\n", name, times)

		time.Sleep(time.Second)
	}
}
