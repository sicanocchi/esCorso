package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("mi chiamo", p.first, "e ho anni:", p.age)
}

func main() {
	p1 := person{
		first: "beppe",
		last:  "rossi",
		age:   17,
	}
	p1.speak()
}
