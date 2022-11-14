package ch07

import "fmt"

func UsingAny() {
	var i any = 20
	var number int = i.(int)
	fmt.Printf("i as int: %d\n", i)
	fmt.Printf("number: %d\n", number)

	i = "some text"
	var message string = i.(string)
	fmt.Printf("i as string: %s\n", i)
	fmt.Printf("message: %s\n", message)
}
