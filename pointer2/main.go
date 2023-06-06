package main

import (
	"fmt"
)

type person struct {
	name string
}

func main() {
	p1 := person{
		name: "beppe",
	}
	p2 := 7
	fmt.Println(p1)
	changeMe(&p1)
	changeMe1(&p2)

	fmt.Println(p1)
}

func changeMe(p *person) {
	(*p).name = "mario"
}

func changeMe1(i *int) {
	*i = 5
}
