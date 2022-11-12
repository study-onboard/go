package ch02

import "fmt"

func VariableAssign() {
  var number1, number2 int64
  number1, number2 = 48, 64
  fmt.Printf("number1: %d, number2: %d\n", number1, number2)

  number1, number2 = number2, number1
  fmt.Printf("number1: %d, number2: %d\n", number1, number2)
}

