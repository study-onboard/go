package ch08

import (
	"flag"
	"log"
)

func UsingFlagPackage() {
	var name string
	var age int
	var married bool

	flag.StringVar(&name, "name", "James", "name")
	flag.IntVar(&age, "age", 6, "age")
	flag.BoolVar(&married, "married", false, "married")
	flag.Parse()

	log.Printf("Name: %s", name)
	log.Printf("Age: %d", age)
	log.Printf("Married: %t", married)
}
