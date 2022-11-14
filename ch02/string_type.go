package ch02

import (
	"fmt"
	"unicode/utf8"
)

func StringType() {
	// compare
	message1 := "show me the money"
	message2 := "show me the money"
	if message1 == message2 {
		fmt.Println("message1 equals message2")
	} else {
		fmt.Println("message1 not equals message2")
	}

	// join
	slogan := "Make Love, No War!!"
	prefix := "Slogan: "
	fmt.Println(prefix + slogan)

	// string length
	words := "我们"
	fmt.Printf("len(words): %d, RuneCountInString(words): %d\n", len(words), utf8.RuneCountInString(words))

	// multiple lines string
	json := `
  {
    "name": "Kut Zhang",
    "position", "IT Manager"
  }
  `
	fmt.Printf("json: %s\n", json)

	// iterial string
	content := "我们是Supper Start!"
	fmt.Println(content)

	fmt.Println("len mode: ")
	for i := 0; i < len(content); i++ {
		fmt.Printf("%d: %c(%d)\n", i, content[i], content[i])
	}

	fmt.Println("range mode: ")
	for i, word := range content {
		fmt.Printf("%d: %c(%d)\n", i, word, word)
	}

	var someText string
	fmt.Printf("Initialized string value is: `%s`\n", someText)

	fmt.Println("-------------------------------------------------------------------")
	message1 = "abcd1"
	message2 = "abcd2"
	fmt.Printf("message 1: %s\n", message1)
	fmt.Printf("message 2: %s\n", message2)

	switch {
	case message1 > message2:
		fmt.Println("message1 > message2")

	case message1 == message2:
		fmt.Println("message1 == message2")

	default:
		fmt.Println("message1 < message2")
	}
}
