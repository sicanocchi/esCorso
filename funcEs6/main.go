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

func main() {
	s1 := square{
		lato: 2,
	}
	c1 := circle{
		raggio: 3,
	}

	fmt.Println(s1.area())
	fmt.Println(c1.area())

	func() {
		fmt.Println("anonymous func")
	}()

}
