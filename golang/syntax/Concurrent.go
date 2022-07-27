package main

import (
	"fmt"
)

func main() {

	go print()
	select {}
}
func print() {
	for _, v := range "HelloWorld" {
		fmt.Println(string(v))
	}
}
