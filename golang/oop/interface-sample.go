package main

import "fmt"

type light interface {
	TurnOn()
	TurnOff()
}

type CarLight struct{}

func (carLigth CarLight) TurnOn() {
	fmt.Println("trun on")
}
func (carLight CarLight) TurnOff() {
	fmt.Println("turn off")
}
func main() {
	var carLight CarLight
	fmt.Println("light")
	carLight.TurnOn()
	carLight.TurnOff()

}
