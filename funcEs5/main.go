package main

import (
	"fmt"
	"math"
)

type square struct {
	lato float64
}

type circle struct {
	raggio float64
}

func (s square) area() float64 {
	return s.lato * s.lato
}

func (c circle) area() float64 {
	return math.Pi * (c.raggio) * (c.raggio)
}

type shape interface {
	area() float64
}

func info(h shape) {
	fmt.Println(h.area())
}

func main() {
	s1 := square{
		lato: 2,
	}
	c1 := circle{
		raggio: 3,
	}

	info(s1)
	info(c1)
}
