package ch06

import (
	"encoding/json"
	"fmt"
)

// references
// ---------------------------------------------------------------
// https://gobyexample.com/json

func UsingJSON() {
	book := &Book{
		Name:  "Go language",
		Price: 104.50,
		Author: Person{
			Name: "Kut Zhang",
			Age:  105,
		},
	}

	jsonText, _ := json.Marshal(book)
	fmt.Printf("%s\n", string(jsonText))

	bookClone := new(Book)
	if err := json.Unmarshal(jsonText, bookClone); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", bookClone)
}

// --- types -----------------------------------------------------
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Book struct {
	Name   string  `json:"name"`
	Author Person  `json:"person"`
	Price  float32 `json:"price"`
}
