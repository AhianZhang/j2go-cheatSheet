package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string, 1)
	go func() {
		channel <- "add"
	}()
	go func() {
		val := <-channel
		fmt.Println(val)
	}()
	time.Sleep(time.Second * 2)
}
