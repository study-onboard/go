package ch02

import "fmt"

func TypeAlias() {
  type Moon = string
  var message Moon = "show me the money"
  fmt.Printf("message: %s\n", message)
}
