package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 打开文件
	f, err := ioutil.ReadFile("file path")
	if err != nil {
		println(err)
		os.Exit(1)
	}
	fmt.Println(string(f))
}
