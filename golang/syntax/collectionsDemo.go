package main

import "fmt"

func main() {
	stack := make([]int, 0)
	stack = append(stack, 1)
	stack = append(stack, 2)
	for len(stack) > 0 {
		topEle := stack[len(stack)-1] // inspect
		fmt.Println("top element is: ", topEle)
		stack = stack[:len(stack)-1] // pop
		fmt.Println("current statck size: ", len(stack))

	}
	stack = append(stack, 3)
	fmt.Println(len(stack))
	stack = stack[:0] // flush
	fmt.Println(len(stack))
}
