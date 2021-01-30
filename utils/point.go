package utils

import "math"

type Point struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

// Retorna la distancia entre dos puntos por pitagoras
func (p Point) GetDistance(p2 Point) float64 {
	disX := math.Pow(p.X-p2.X, 2)
	disY := math.Pow(p.Y-p2.Y, 2)

	return math.Sqrt(disX + disY)
}
