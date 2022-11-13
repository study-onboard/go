package ch02

import "fmt"

func TypeAlias() {
  type MooX = string
  type MooY string
  var info MooX = "show me the money"
  var error MooY = "how do you turn this on"

  // will print `string`
  fmt.Printf("type of info: %T\n", info)

  // will print `ch02.MooY`
  fmt.Printf("type of error: %T\n", error)

  fmt.Printf("info: %s\n", info)
  fmt.Printf("error: %s\n", error)
}
