package main

import "fmt"

func main() {

	print("hello")
	print(123)
}

func print[D any](data D) {
	fmt.Print("your data is: ", data)
}
