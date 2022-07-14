package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 声明方式 1
	var str1 string
	var int1 int
	// 声明方式 2
	str2 := ""
	// 常量
	const str3 = "con"

	fmt.Println(str1)
	fmt.Println(int1)
	fmt.Println(str2)
	fmt.Println(str3)

	// 循环控制
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// range
	str4 := "abc"
	for idx, val := range str4 {
		fmt.Println("index:", idx, "val:", val)
	}
	// array list
	list1 := make([]string, 2)
	list1 = append(list1, "a", "b", "c")
	fmt.Println("list1 val is", list1, "len is ", len(list1), " capicity is ", cap(list1))

	// map
	maps := make(map[int]string)
	maps[1] = "a"
	val, ok := maps[1]
	if ok {
		fmt.Println(val)
	}
	// set
	sets := make(map[string]struct{})
	sets["a"] = struct{}{}
	if _, ok := sets["a"]; ok {
		fmt.Println("element exist")
	}
	// 优化 set
	type Void struct{}
	var x Void
	set := make(map[string]Void)
	set["a"] = x
	if _, ok := sets["a"]; ok {
		fmt.Println("element exist")
	}
	// var arg string
	// arg = os.Args[0]
	fmt.Println(os.Args[1:])
	// 命令行
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		fmt.Println(input.Text())
	}

	arr := [10]int{}
	fmt.Println(len(arr))
}
