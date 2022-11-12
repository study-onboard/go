package ch02

import "fmt"

func IntType() {
  var (
    int16Number int16
    int32Number int32
    int64Number int64

    uint16Number uint16
    uint32Number uint32
    uint64Number uint64
  )

  int16Number = -32
  int32Number = -65535
  int64Number = -3265535

  uint16Number = 32
  uint32Number = 65535
  uint64Number = 3265535

  fmt.Printf("int16Number: %d\n", int16Number)
  fmt.Printf("int32Number: %d\n", int32Number)
  fmt.Printf("int64Number: %d\n", int64Number)
  fmt.Printf("uint16Number: %d\n", uint16Number)
  fmt.Printf("uint32Number: %d\n", uint32Number)
  fmt.Printf("uint64Number: %d\n", uint64Number)
}
