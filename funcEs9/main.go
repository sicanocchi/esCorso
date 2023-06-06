package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func substract(a int, b int) int {
	return a - b
}

func operation(a, b int, f func(int, int) int) int {
	return f(a, b)
}
func main() {

	x := operation(4, 5, add)
	y := operation(7, 5, substract)

	fmt.Println(x)
	fmt.Println(y)

}
