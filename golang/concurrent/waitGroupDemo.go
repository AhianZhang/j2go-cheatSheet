package main

import (
	"fmt"
	"sync"
)

var num = 0

// 计数器
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1) // 计数器加 1
		go add(&wg)
	}
	wg.Wait() // 等待计数器归零
	fmt.Println(num)
}
func add(wg *sync.WaitGroup) {
	num++
	wg.Done() // 计数器减 1
}
