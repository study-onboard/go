package ch05

import (
	"fmt"
	"github.com/google/uuid"
)

func UsingClosure() {
	simplyUse()
	fmt.Println("-------------------------------------------------------------")
	complexUse()
}

func buildIdGenerator(prefix string) func() string {
	return func() string {
		var subfix = uuid.New()
		return fmt.Sprintf("%s-%s", prefix, subfix.String())
	}
}

func complexUse() {
	generator := buildIdGenerator("MDC")
	id := generator()
	fmt.Printf("id: %s\n", id)

	id = generator()
	fmt.Printf("id: %s\n", id)
}

func simplyUse() {
	var number int32 = 32
	fmt.Printf("Before number changed, number is %d\n", number)

	var changeNumber = func() {
		number = 64
	}

	changeNumber()
	fmt.Printf("After number changed, number is %d\n", number)
}
