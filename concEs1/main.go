package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()
	wg.Wait()
}

func foo() (int, string) {
	wg.Done()
	return rand.Intn(9), "pluto"
}

func bar() (int, string) {
	return rand.Intn(9), "pippo"
}
