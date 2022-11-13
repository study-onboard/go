package ch05

import "fmt"

func ImplementInterface() {
	var animal animal = &human{
		name: "Kut Zhang",
	}
	animal.run()
}

type animal interface {
	run()
}

type human struct {
	name string
}

func (human *human) run() {
	fmt.Printf("%s begin run.\n", human.name)
	fmt.Printf("%s runing........\n", human.name)
	fmt.Printf("%s run complete.\n", human.name)
}
