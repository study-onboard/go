package ch02

// References
// --------------------------------------
// https://www.golangprograms.com/how-to-convert-string-to-integer-type-in-go.html
// 

import (
	"fmt"
	"math"
	"strconv"
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


  var intString string = "193383"
  convertStringToInt(intString)
  intString = "ZD0091"
  convertStringToInt(intString)
}

func convertStringToInt(intString string) {
  intNumber, error := strconv.Atoi(intString)
  if error != nil {
    fmt.Printf("Type of error: %T\n", error)
    fmt.Printf("Can't not convert '%s' to number: %s\n", intString, error)
  } else {
    fmt.Printf("'%s' converted to int number: %d\n", intString, intNumber)
  }
}
