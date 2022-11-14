package ch07

import "fmt"

func CombinationInterface() {
	var programer Programer = newCodeFarmer("Grow Liu")
	programer.fuck()
	programer.juan()
}

type Phper interface {
	fuck()
}

type Javaer interface {
	juan()
}

type Programer interface {
	Phper
	Javaer
}

type CodeFarmer struct {
	name string
}

func newCodeFarmer(name string) *CodeFarmer {
	return &CodeFarmer{
		name: name,
	}
}

func (farmer *CodeFarmer) fuck() {
	fmt.Println("Fuck you")
}

func (farmer *CodeFarmer) juan() {
	fmt.Println("Juan si you")
}
