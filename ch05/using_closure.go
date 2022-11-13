package ch05

import (
	"fmt"
	"github.com/google/uuid"
)

func UsingClosure() {
	generator := buildIdGenerator("MDC")
	id := generator()
	fmt.Printf("id: %s\n", id)

	id = generator()
	fmt.Printf("id: %s\n", id)
}

func buildIdGenerator(prefix string) func() string {
	return func() string {
		var subfix = uuid.New()
		return fmt.Sprintf("%s-%s", prefix, subfix.String())
	}
}
