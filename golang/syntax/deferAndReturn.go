package main

// return 在 defer 前执行
// 如果 defer 在 panic 之前能执行，panic 之后代码会中断
import "fmt"

func main() {
	deferReturnPanicFun()

}
func deferReturnPanicFun() int {

	defer deferFuc()
	panicFun()
	return returnFuc()

}
func deferFuc() int {
	fmt.Println("defer")
	return 0
}
func returnFuc() int {
	fmt.Println("return")
	return 0
}
func panicFun() {
	panic("panic")
}
