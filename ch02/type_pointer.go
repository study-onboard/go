package ch02

import "fmt"

func TypePointer() {
  // known what is pointer
  var number int = 32
  var message string = "show me the money"
  fmt.Printf("number pointer: %p, message pointer: %p\n", &number, &message)

  // print information of pointer
  var messagePointer *string = &message
  fmt.Printf("message pointer type: %T\n", messagePointer)
  fmt.Printf("message pointer value: %p\n", messagePointer)
  fmt.Printf("message value by pointer: %s\n", *messagePointer)

  // modify variable value by pointer
  fmt.Printf("number: %d\n", number)
  var numberPointer *int = &number
  *numberPointer = 64
  fmt.Printf("number: %d\n", number)
  
  fmt.Printf("message: %s\n", message)
  *messagePointer = "how do you turn this on?"
  fmt.Printf("message: %s\n", message)

  // create pointer by new
  var contentPtr = new(string)
  *contentPtr = "Good good study, Day day up!!"
  fmt.Printf("content: %s\n", *contentPtr)
}
