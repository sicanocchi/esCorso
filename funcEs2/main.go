package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(foo(2, 3, 4))
	x := []int{2, 3, 4, 5, 6}
	fmt.Println(foo(x...))
	fmt.Println(bar())
}

func foo(x ...int) int {
	fmt.Println(x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar() (int, string) {
	return rand.Intn(9), "pippo"
}
