package main

import (
	"fmt"
	"time"
)

func TimedFunction(fn func()) {
	start := time.Now()
	fn()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func MyFunc() {
	time.Sleep(2 * time.Second)
	fmt.Println("my func complete")
}

func main() {
	TimedFunction(MyFunc)
}
