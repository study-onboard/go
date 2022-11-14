package ch05

import "fmt"

func UsingPanic() {
	defer fmt.Printf("show me the money\n")
	panic("Just a panic test")
}
