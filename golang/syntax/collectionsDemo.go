package main

import (
	"fmt"
	"sort"
)

func main() {
	stack()
}

func stack() {
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

func sliceSort() {
	sli := make([]int, 10)
	sort.Ints(sli)
	sort.Slice(sli, func(a, b int) bool { return sli[a] < sli[b] })
	sort.SliceStable(sli, func(a, b int) bool { return sli[a] < sli[b] }) // 如果元素相同保持当前位置不变
}
