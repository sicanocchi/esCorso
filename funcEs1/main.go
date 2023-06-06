package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return rand.Intn(5)
}

func bar() (int, string) {
	return rand.Intn(9), "pippo"
}
