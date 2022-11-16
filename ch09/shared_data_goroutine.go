package ch09

import (
	"fmt"
	"runtime"
	"sync"
)

func SharedDataGoroutine() {
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go count(i, lock)
	}

	for {
		lock.Lock()
		c := counter
		lock.Unlock()

		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
}

var counter int64 = 0

func count(number int, lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Printf("%3d: %d\n", number, counter)
	lock.Unlock()
}
