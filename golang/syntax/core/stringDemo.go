package main

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

func main() {
	// rune alias int32
	str := "abc"
	for _, v := range str {
		fmt.Println(v, reflect.TypeOf(v)) //rune
	}
	for i := range str {
		fmt.Println(str[i], reflect.TypeOf(str[i])) //byte
	}
	r := []rune(str)
	for _, v := range r {
		fmt.Println(v, reflect.TypeOf(v))
	}
	fmt.Println("IsLetter", unicode.IsLetter('a'))
	fmt.Println("IsDigit", unicode.IsDigit('1'))
	fmt.Println("IsLower", unicode.IsLower('A'))
	fmt.Println("IsNumber", unicode.IsNumber('1'))
	fmt.Println(reflect.TypeOf("1123"[1])) // byte
	strs := strings.Split("a,b,c", ",")
	fmt.Println(strs, reflect.TypeOf(strs))
	var s string
	strings.HasPrefix(s, "prefix")
	strings.HasSuffix(s, "suffix")
}
