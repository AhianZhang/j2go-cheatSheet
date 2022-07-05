package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type user struct {
	Name string
	Age  int
}

func main() {
	jsonDecode()
}

func jsonDecode() {
	user := &user{Name: "ahian", Age: 18}
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Errorf("could not marshal json: %s\n", err)
		return
	}
	fmt.Println(string(b))
}
func readFile() {
	// 打开文件
	f, err := ioutil.ReadFile("file path")
	if err != nil {
		println(err)
		os.Exit(1)
	}
	fmt.Println(string(f))
}
