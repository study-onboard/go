package ch03

import (
	"fmt"
	"sync"
)

func UsingSyncMap() {
	var masaga sync.Map

	masaga.Store("name", "Kut Zhang")
	masaga.Store("position", "IT Manager")

	masaga.Range(func(key, value any) bool {
		fmt.Printf("%s: %s\n", key, value)
		return true
	})
}
