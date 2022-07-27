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
	//jsonDecode()
	readFile2()
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
func readFile1() {
	// 打开文件
	f, err := ioutil.ReadFile("file path")
	if err != nil {
		println(err)
		os.Exit(1)
	}
	fmt.Println(string(f))
}
func readFile2() {
	// 打开文件
	file, err := os.Open("file path")
	if err != nil {
		return
	}
	defer file.Close()
	buffer := make([]byte, 1024)
	file.Read(buffer)
	fmt.Println(string(buffer))
}
