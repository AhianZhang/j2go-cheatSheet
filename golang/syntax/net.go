package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := os.Args[1]
	r, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {

	}
	// 打印
	fmt.Println(string(b))
	// 复制到指定位置
	// io.Copy()

}
