package ch05

import "fmt"

func UsingRecover() {
	usingRecoverProcess()
	fmt.Println("How do you turn this on?")
}

func usingRecoverProcess() {
	defer func() {
		err := recover()
		fmt.Printf("Recover from: %s\n", err)
	}()

	panic("Show me the money.")
	fmt.Println("After show me the money")
}
