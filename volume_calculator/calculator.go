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

func PrintSphereVolume(radius float64) {

	sphere := Sphere{Radius: radius}

	fmt.Printf("Volume of Sphere: %.2f\n", sphere.Volume())

}

func PrintCubeVolume(length float64) {

	cube := Cube{Length: length}

	fmt.Printf("Volume of Cube: %.2f\n", cube.Volume())

}

func main() {

	PrintCubeVolume(3.0)

	PrintSphereVolume(10.0)

}
