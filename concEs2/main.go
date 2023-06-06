package main

import (
	"fmt"
)

type person struct {
	nome string
	age  int
}

func (p *person) speak() {
	fmt.Println("mi chiamo", (*p).nome)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		nome: "mario",
		age:  20,
	}
	saySomething(&p1)
}
