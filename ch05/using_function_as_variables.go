package ch05

import "fmt"

func UsingFunctionAsVariables() {
	var simpFunc func() = simpleFunction
	simpFunc()

	var parametersFunc func(string, string) = parametersFunction
	parametersFunc("shutdown the server", "MDC")

	calculateSum(25, 5, callback)
}

func simplyFunction() {
	fmt.Println("Simple Function")
}

func parametersFunction(message string, prefix string) {
	fmt.Printf("sending message: %s - %s\n", prefix, message)
}

func calculateSum(base int32, times int32, callback func(int32)) {
	sum := base
	var i int32
	for i = 0; i < times; i++ {
		sum += i
	}
	callback(sum)
}

func callback(result int32) {
	fmt.Printf("result: %d\n", result)
}
