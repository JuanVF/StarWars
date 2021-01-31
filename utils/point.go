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

// Parsea una lista de flotantes con formato {x, y} a una lista de puntos
func (p Point) Parse(points []float64) []Point {
	// El formato no es correcto
	if len(points)%2 != 0 {
		return nil
	}

	rst := []Point{}

	for i := 0; i < len(points); i += 2 {
		rst = append(rst, Point{
			X: points[i],
			Y: points[i+1],
		})
	}

	return rst
}
