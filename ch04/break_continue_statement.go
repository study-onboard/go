package ch04

import "fmt"

func BreakContinueStatement() {
	breakStatement()
	fmt.Println("---------------------------------------------------------------------------------")
	continueStatement()
}

func breakStatement() {
	for i := 0; i < 5; i++ {
	ILoop:
		for j := 0; j < 5; j++ {
			for k := 0; k < 5; k++ {
				if k > 3 {
					fmt.Printf("Break to ILoop: i: %d, j: %d, k: %d\n", i, j, k)
					break ILoop
				}
				fmt.Printf("i: %d, j: %d, k: %d\n", i, j, k)
			}
		}
	}
}

func continueStatement() {
	for i := 0; i < 5; i++ {
	ILoop:
		for j := 0; j < 5; j++ {
			for k := 0; k < 5; k++ {
				if k > 3 {
					fmt.Printf("Continue to ILoop: i: %d, j: %d, k: %d\n", i, j, k)
					continue ILoop
				}
				fmt.Printf("i: %d, j: %d, k: %d\n", i, j, k)
			}
		}
	}
}
