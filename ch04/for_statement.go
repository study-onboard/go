package ch04

import (
	"fmt"
	"strings"
)

func ForStatement() {
	// simple use
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Printf("sum + %d: %d\n", i, sum)
	}

	fmt.Println("-------------------------------------------")

	// loop 1
	sum = 0
	for {
		sum += 1
		sum += sum
		if sum > 100 {
			break
		}
	}
	fmt.Printf("sum: %d\n", sum)
	fmt.Println("-------------------------------------------")

	// loop 2
	sum = 0
	for ; sum < 100; sum += 1 {
		sum += sum
	}
	fmt.Printf("sum: %d\n", sum)
	fmt.Println("-------------------------------------------")

	// loop 3
	sum = 0
	for sum < 100 {
		sum += 1
		sum += sum
	}
	fmt.Printf("sum: %d\n", sum)
	fmt.Println("-------------------------------------------")

	// rune loop for string
	message := "美丽国的国运没了, USE will dead"
	for index, char := range message {
		fmt.Printf("%d: %c\n", index, char)
	}
	guoIndex := strings.IndexAny(message, "国")
	fmt.Printf("Guo index: %d\n", guoIndex)
	message = message[guoIndex:]
	fmt.Printf("message: %s\n", message)
}
