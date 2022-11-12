package ch02

import "fmt"

func AnonymousVariables() {
  message1, _ := generateMessages()
  _, message2 := generateMessages()

  fmt.Printf("message 1: %s, message 2: %s\n", message1, message2)
}

func generateMessages() (string, string) {
  return "Show me the money", "How do you turn this on"
}
