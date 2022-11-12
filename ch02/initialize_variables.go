package ch02 

import "fmt"

func InitializeVariables() {
  var name string = "kut"
  var position = "manager"
  age := 42
  fmt.Printf("name: %s, position: %s, age: %d\n", name, position, age)

  message, times := generateMessage()
  fmt.Printf("message: %s, times: %d\n", message, times)

  message, _ = generateMessage()
  fmt.Printf("message: %s, times: %d\n", message, times)
}


func generateMessage() (string, int) {
  return "Show me the money", 44
}
