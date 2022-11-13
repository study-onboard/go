package ch05

import "fmt"

func DefineFunction() {
	simpleFunction()
	printMessage(simpleReturnFunction(true))
	printMessage(simpleReturnFunction(false))

	times, message := multipleReturnFunction()
	fmt.Printf("times: %d, message: %s\n", times, message)

	times, message = namedMultipleReturnFunction()
	fmt.Printf("times: %d, message: %s\n", times, message)

	var sum int
	ptrParameterFunction(2, 3, &sum)
	fmt.Printf("sum: %d\n", sum)
}

func simpleFunction() {
	fmt.Println("Simple function")
}

func simpleReturnFunction(checked bool) *string {
	if checked {
		message := "show me the money"
		return &message
	} else {
		return nil
	}
}

func printMessage(messagePtr *string) {
	if messagePtr != nil {
		fmt.Printf("message: %s\n", *messagePtr)
	} else {
		fmt.Println("no message")
	}
}

func multipleReturnFunction() (int, string) {
	return 10, "how do you turn this on?"
}

func namedMultipleReturnFunction() (times int, message string) {
	times = 20
	message = "good good study, day day up!"
	return
}

func ptrParameterFunction(a int, b int, sum *int) {
	*sum = a + b
}
