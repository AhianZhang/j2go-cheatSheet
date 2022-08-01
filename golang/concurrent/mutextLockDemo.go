package main

import (
	"fmt"
	"sync"
	"time"
)

var n = 0

func main() {
	var lock sync.Mutex
	for i := 0; i < 10000; i++ {
		go addByLock(&lock)
	}
	time.Sleep(10)
	fmt.Println(n)
}
func addByLock(lock *sync.Mutex) {
	lock.Lock()
	n++
	lock.Unlock()
}
