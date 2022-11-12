package ch02

import (
	"flag"
	"fmt"
)

func UsingArgs() {
  envPtr := flag.String("env", "", "application environment, such as dev / sit / qa / prod")
  flag.Parse()
  fmt.Printf("Application environment: %s\n", *envPtr)
}
