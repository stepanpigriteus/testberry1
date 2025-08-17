package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками на плоскости. Точки представлены в виде структуры Point с инкапсулированными (приватными) полями x, y (типа float64) и конструктором. Расстояние рассчитывается по формуле между координатами двух точек.
// Подсказка: используйте функцию-конструктор NewPoint(x, y), Point и метод Distance(other Point) float64.

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p *Point) Distance(other Point) float64 {
	return math.Sqrt(math.Pow(other.x-p.x, 2) + math.Pow(other.y-p.y, 2))
}

func main() {
	p1 := NewPoint(1.0, 2.0)
	p2 := NewPoint(4.0, 6.0)
	fmt.Printf("Distance between p1 and p2: %.2f\n", p1.Distance(p2))
}
