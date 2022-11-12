package ch02

import (
	"fmt"
	"math"
)

func TypeConvert() {
  var int16Number int16 = 16
  var int64Number int64 = 933668
  int16Number = int16(int64Number)
  fmt.Printf("int16Number: %d\n", int16Number)

  fmt.Printf("int8 range: %d, %d\n", math.MinInt8, math.MaxInt8)
  fmt.Printf("int16 range: %d, %d\n", math.MinInt16, math.MaxInt16)
  fmt.Printf("int32 range: %d, %d\n", math.MinInt32, math.MaxInt32)
  fmt.Printf("int64 range: %d, %d\n", math.MinInt64, math.MaxInt64)

  fmt.Printf("PI: %f\n", math.Pi)
}
