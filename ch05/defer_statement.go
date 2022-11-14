package ch05

import "fmt"

func DeferStatement() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	defer fmt.Println("defer 4")
	defer fmt.Println("defer 5")
	defer fmt.Println("defer 6")
}
