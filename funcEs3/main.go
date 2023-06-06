package main

import (
	"fmt"
)

func main() {
	fmt.Println("----")
	defer bar()
	x := []int{2, 3, 4, 5, 6}
	i := foo(x...)
	fmt.Println(i)

}

func foo(x ...int) int {
	fmt.Println(x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar() {
	fmt.Println("sono il defer")
}
