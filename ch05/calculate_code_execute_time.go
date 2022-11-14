package ch05

import (
	"fmt"
	"time"
)

func CalculateCodeExecuteTime() {
	start := time.Now()

	sum := int64(0)
	for i := int64(0); i < 1000000000; i++ {
		sum += i
	}
	fmt.Printf("sum: %d\n", sum)

	elapsed := time.Since(start)
	fmt.Printf("Execute time: %s\n", elapsed)
}
