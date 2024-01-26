package main

import (
	"math"
	"testing"
)

func TestSphereVolume(t *testing.T) {
	tests := []struct {
		radius         float64
		expectedVolume float64
	}{
		{5.0, (4.0 / 3.0 * math.Pi * math.Pow(5.0, 3))},
		{10.0, (4.0 / 3.0 * math.Pi * math.Pow(10.0, 3))},
		{6.7, (4.0 / 3.0 * math.Pi * math.Pow(6.7, 3))},
		{1.4, (4.0 / 3.0 * math.Pi * math.Pow(1.4, 3))},
	}

	for _, tt := range tests {
		t.Run("SphereVolume", func(t *testing.T) {
			sphere := Sphere{Radius: tt.radius}
			result := sphere.Volume()
			if result != tt.expectedVolume {
				t.Errorf("error: expected %.2f, got %.2f \n", tt.expectedVolume, result)
			}
		})
	}
}

func TestCubeVolume(t *testing.T) {
	tests := []struct {
		length         float64
		expectedVolume float64
	}{
		{3.0, math.Pow(3.0, 3)},
		{5.0, math.Pow(5.0, 3)},
		{6.6, math.Pow(6.6, 3)},
		{5.55, math.Pow(5.55, 3)},
	}

	for _, tt := range tests {
		t.Run("CubeVolume", func(t *testing.T) {
			cube := Cube{Length: tt.length}
			result := cube.Volume()
			if result != tt.expectedVolume {
				t.Errorf("error: expected %.2f, got %.2f \n", tt.expectedVolume, result)
			}
		})
	}
}
