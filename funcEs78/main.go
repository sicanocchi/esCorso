package main

import "fmt"

func speak() int {
	return 42
}

func main() {
	var n = speak()
	fmt.Println(n)

	f := func() {
		for i := 0; i < 4; i++ {
			fmt.Println(i)
		}
	}

	f()

}
