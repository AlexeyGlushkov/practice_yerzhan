package main

import (
	"fmt"
	"math"
)

type Figure interface {
	Volume()
}

type Sphere struct {
	Radius float64
}

func (s Sphere) Volume() float64 {
	return (4.0 / 3.0) * math.Pi * math.Pow(s.Radius, 3)
}

type Cube struct {
	Length float64
}

func (c Cube) Volume() float64 {
	return math.Pow(c.Length, 3)
}

func main() {
	sphere := Sphere{2.0}
	cube := Cube{3.0}

	fmt.Println(sphere, cube)
}
